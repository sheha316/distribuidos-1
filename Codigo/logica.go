package main
import (
  "log"
  "net"
  "google.golang.org/grpc"
  "github.com/sheha316/distribuidos-1/Codigo/comms"
  "golang.org/x/net/context"
  "os"
  "io"
  "encoding/csv"
  "fmt"
  "strconv"
  "bufio"
  "time"
  "github.com/streadway/amqp"
)

type paquete struct{
  Id string
  Seguimiento int
  Tipo string
  Valor int
  Intentos int
  Estado string
}

type envio struct{
  Id_paquete string
  Estado string
  Id_camion string
  Seguimiento int
  Tipo string
  Valor int
  Intentos int
  Uso string
}

type Server struct{
  envios_s [6]envio
  candado bool
}

func find_file(nombre string,tipo string)(string){
  prefijo:="0"
  switch tipo {
  case "retail":
    prefijo="1"
  default:
    prefijo="2"
  }
  seguimento:=0
  file,erros:=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  }
  file.Close()
  log.Printf("Seguimiento entregado: %s", prefijo+strconv.Itoa(seguimento))
  return prefijo+strconv.Itoa(seguimento)
}

func registro_logico_pymes(tipo string,request *comms.Request_CrearOrdenPyme)(int){
  seguimento:=find_file("-",tipo)
  file,erros:=os.Create("../storage/logica/"+seguimento+".csv")
  if erros!=nil{
    fmt.Println(erros)
  }
  writer:=csv.NewWriter(file)
  t := time.Now()
  seguimentoint,_:=strconv.Atoi(seguimento)
  var guardar = [][]string{
    {t.String(),request.Id,tipo,request.Producto,strconv.Itoa(int(request.Valor)),request.Tienda,request.Destino,seguimento},
  }
  erros=writer.WriteAll(guardar)
  file.Close()
  return seguimentoint
}

func registro_logico_retail(tipo string,request *comms.Request_CrearOrdenRetail)(int){
  seguimento:=find_file("-",tipo)
  file,erros:=os.Create("../storage/logica/"+seguimento+".csv")
  if erros!=nil{
    fmt.Println(erros)
  }
  writer:=csv.NewWriter(file)
  t := time.Now()
  seguimentoint,_:=strconv.Atoi(seguimento)
  var guardar = [][]string{
    {t.String(),request.Id,tipo,request.Producto,strconv.Itoa(int(request.Valor)),request.Tienda,request.Destino,seguimento},
  }
  erros=writer.WriteAll(guardar)
  file.Close()
  return seguimentoint
}

func registro_paquete_pymes(request *comms.Request_CrearOrdenPyme,seguimento int){
  f, err := os.OpenFile("../storage/logica/pymes.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    log.Printf("registro_paquete_pymes")
    log.Fatal(err)
  }
  defer f.Close()
  tipo:="normal"
  if(int(request.Prioritario)==1){
    tipo="prioritario"
  }
  var data [][]string
  data = append(data, []string{request.Id,strconv.Itoa(seguimento),tipo,strconv.Itoa(int(request.Valor)),"0","En bodega"})
  w := csv.NewWriter(f)
  w.WriteAll(data)
  if err := w.Error(); err != nil {
    log.Printf("registro_paquete_pymes")
    log.Fatal(err)
  }
}

func registro_paquete_retail(request *comms.Request_CrearOrdenRetail,seguimento int){
  f, err := os.OpenFile("../storage/logica/retail.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    log.Printf("registro_paquete_retail")
    log.Fatal(err)
  }
  defer f.Close()
  var data [][]string
  data = append(data, []string{request.Id,strconv.Itoa(seguimento),"retail",strconv.Itoa(int(request.Valor)),"0","En bodega"})
  w := csv.NewWriter(f)
  w.WriteAll(data)
  if err := w.Error(); err != nil {
    log.Fatal(err)
  }
}

func (s *Server) CrearOrdenPyme(ctx context.Context, request *comms.Request_CrearOrdenPyme) (*comms.Response_CrearOrden, error) {
  log.Printf("Receive message %+v", request)
  for s.candado{}
  s.candado=true
  seguimento:=registro_logico_pymes("pyme",request)
  registro_paquete_pymes(request,seguimento)
  s.candado=false
  return &comms.Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *comms.Request_CrearOrdenRetail) (*comms.Response_CrearOrden, error) {
  log.Printf("Receive message %+v", request)
  for s.candado{}
  s.candado=true
  seguimento:=registro_logico_retail("retail",request)
  registro_paquete_retail(request,seguimento)
  s.candado=false
  return &comms.Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, request *comms.Request_Seguimiento) (*comms.Response_Seguimiento, error) {
  log.Printf("Receive message %d", request.Seguimiento)
  for s.candado{}
  s.candado=true
  for i:=0;i<6;i++{
    if(s.envios_s[i].Seguimiento==int(request.Seguimiento)){
      log.Printf("desde memoria :)")
      s.candado=false
      return &comms.Response_Seguimiento{Estado: s.envios_s[i].Estado}, nil
    }
  }
  aux:=strconv.Itoa(int(request.Seguimiento))
  csvFile,error:=os.Open("../storage/logica/"+aux+".csv")
  if error !=nil{
    s.candado=false
    return &comms.Response_Seguimiento{Estado: "Paquete no existe"}, nil
  }
  reader := csv.NewReader(bufio.NewReader(csvFile))
  line,_ :=reader.Read()
  id:=line[1]
  file:="../storage/logica/pymes.csv"
  if(line[2]=="retail"){
    file="../storage/logica/retail.csv"
  }
  csvFile.Close()
  csvFile,_=os.Open(file)
  reader = csv.NewReader(bufio.NewReader(csvFile))
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Fatal(error)
    }
    if(line[0]==id){
      csvFile.Close()
      s.candado=false
      return &comms.Response_Seguimiento{Estado: line[5]}, nil
    }
  }
  csvFile.Close()
  s.candado=false
  return &comms.Response_Seguimiento{Estado: "Esto no deberia suceder :)"}, nil
}

func (s *Server) SolicitarPaquete(ctx context.Context, request *comms.Request_SolicitarPaquete) (*comms.Response_SolicitarPaquete, error) {
  log.Printf("Receive message %+v", request)
  for s.candado{}
  s.candado=true
  x:=&paquete{Valor: -1,}
  switch request.Tipo {
  case "retail":
    LFP_R(x)
    if(x.Valor==-1){
      LFP_P(x,"prioritario")
    }
  default:
    LFP_P(x,"prioritario")
    if(int(x.Valor)==-1){
      LFP_P(x,"normal")
    }
  }
  if(x.Valor==-1){
    s.candado=false
    return &comms.Response_SolicitarPaquete{Id:"0",Seguimiento:int32(0),Tipo:"pablo",Valor:int32(x.Valor),Tienda:"--",Destino:"--",}, nil

  }
  aux:=strconv.Itoa(int(x.Seguimiento))
  csvFile,_:=os.Open("../storage/logica/"+aux+".csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  line,_:=reader.Read()
  csvFile.Close()
  Updater("../storage/logica/"+aux+".csv","En camino","0")

  for i:=0;i<6;i++{
    if(s.envios_s[i].Uso=="0"){
      s.envios_s[i].Id_paquete=x.Id
      s.envios_s[i].Estado="En camino"
      s.envios_s[i].Id_camion=request.Id
      s.envios_s[i].Seguimiento=int(x.Seguimiento)
      s.envios_s[i].Tipo=x.Tipo
      s.envios_s[i].Valor=x.Valor
      s.envios_s[i].Intentos=0
      s.envios_s[i].Uso="1"
      break
    }
  }/**/
  s.candado=false
  return &comms.Response_SolicitarPaquete{Id:x.Id,Seguimiento:int32(x.Seguimiento),Tipo:x.Tipo,Valor:int32(x.Valor),Tienda:line[5],Destino:line[6],}, nil
}

func Updater(n_file string,estado string,intentos_u string){
  csvfile ,_:= os.Open(n_file)
  reader := csv.NewReader(bufio.NewReader(csvfile))
  linez,_ :=reader.Read()
  change_id:=linez[1]
  change_tipo:=linez[2]
  csvfile.Close()
  nombrearch:="../storage/logica/retail.csv"
  switch change_tipo{
  case "retail":
    nombrearch="../storage/logica/retail.csv"
  default:
    nombrearch="../storage/logica/pymes.csv"
  }
  var data [][]string
  csvfilez ,_:= os.Open(nombrearch)
  readerz := csv.NewReader(bufio.NewReader(csvfilez))
  for{
    line,error :=readerz.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Printf("updater")
        log.Printf("%+v",data)
        log.Printf("Seguimiento: %s", n_file)
        log.Printf(nombrearch)
        log.Fatal(error)
        continue
    }
    switch line[0] {
      case change_id:
        data = append(data, []string{line[0],line[1],line[2],line[3],intentos_u,estado})
      default:
        data = append(data, []string{line[0],line[1],line[2],line[3],line[4],line[5]})

    }
  }
  csvfilez.Close()
  os.Remove(nombrearch)
  csvfilex ,_:= os.OpenFile(nombrearch, os.O_WRONLY|os.O_CREATE, 0777)
  writer:=csv.NewWriter(csvfilex)
  writer.WriteAll(data)
  csvfilex.Close()
}
func LFP_R(pakete *paquete){
  file,_:=os.Open("../storage/logica/retail.csv")
  reader := csv.NewReader(bufio.NewReader(file))
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Printf("LFP_R")
        log.Fatal(error)
    }
    if(line[5]=="En bodega"){
      aux1,_:=strconv.Atoi(line[1])
      aux2,_:=strconv.Atoi(line[3])
      aux3,_:=strconv.Atoi(line[4])
      pakete.Id=line[0]
      pakete.Seguimiento=aux1
      pakete.Tipo=line[2]
      pakete.Valor=aux2
      pakete.Intentos=aux3
      pakete.Estado="En camino"
      file.Close()
      return
    }
  }
}

func LFP_P(pakete *paquete,p string){
  file,_:=os.Open("../storage/logica/pymes.csv")
  reader := csv.NewReader(bufio.NewReader(file))
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Printf("LFP_P")
        log.Fatal(error)
    }
    if( (line[5]=="En bodega") && (p==line[2])){
      aux1,_:=strconv.Atoi(line[1])
      aux2,_:=strconv.Atoi(line[3])
      aux3,_:=strconv.Atoi(line[4])
      pakete.Id=line[0]
      pakete.Seguimiento=aux1
      pakete.Tipo=line[2]
      pakete.Valor=aux2
      pakete.Intentos=aux3
      pakete.Estado="En camino"
      file.Close()
      return
    }
  }
}


func (s *Server) InformarEstado(ctx context.Context, request *comms.Request_Estado) (*comms.Response_Estado, error) {
  log.Printf("Receive message %+v", request)
  for s.candado{}
  s.candado=true
  for i:=0;i<6;i++{
    if(s.envios_s[i].Id_paquete==request.Id){
      Updater("../storage/logica/"+strconv.Itoa(s.envios_s[i].Seguimiento)+".csv",request.Estado,strconv.Itoa(int(request.Intentos)) )
      tipo_fin:=s.envios_s[i].Tipo
      valor_fin:=s.envios_s[i].Valor
      s.envios_s[i].Uso="0"
      break
    }
  }
  conn, err := amqp.Dial("amqp://test:test@dist96:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"finances", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "{id: "+request.Id+",tipo:"+tipo_fin+",valor:"+strconv.Itoa(valor_fin)+",intentos:"+strconv.Itoa(int(request.Intentos))+",fecha:"+request.Fecha+"}"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
  s.candado=false
  return &comms.Response_Estado{Recibido: "holo"}, nil
}

func (s *Server) LimpiarRegistros(ctx context.Context, request *comms.Dummy) (*comms.Dummy, error){
  for s.candado{}
  s.candado=true
  seguimento:=0
  prefijo:="1"
  file,erros:=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    file.Close()
    os.Remove("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
    seguimento++
    file,erros=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  }
  file.Close()
  prefijo="2"
  seguimento=0
  file,erros=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    file.Close()
    os.Remove("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
    seguimento++
    file,erros=os.Open("../storage/logica/"+prefijo+strconv.Itoa(seguimento)+".csv")
  }
  file.Close()
  os.Remove("../storage/logica/pymes.csv")
  os.Remove("../storage/logica/retail.csv")
  os.Create("../storage/logica/retail.csv")
  os.Create("../storage/logica/pymes.csv")
  s.candado=false
  return &comms.Dummy{Id:"1"}, nil

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
  lis, err := net.Listen("tcp", ":9001")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := Server{}
  for i:=0;i<6;i++{
    s.envios_s[i].Uso="0"
  }
  s.candado=false
  grpcServer := grpc.NewServer()
  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
