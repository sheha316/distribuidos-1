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
  log.Printf("AUXILIAR message %s", prefijo+strconv.Itoa(seguimento))
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

func registro_paquete_pymes_pymes(request *comms.Request_CrearOrdenPyme,seguimento int){
  f, err := os.OpenFile("../storage/logica/pymes.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
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
    log.Fatal(err)
  }
}

func registro_paquete_pymes_retail(request *comms.Request_CrearOrdenRetail,seguimento int){
  f, err := os.OpenFile("../storage/logica/retail.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
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
  log.Printf("Receive message %s", request.Id)
  seguimento:=registro_logico_pymes("pyme",request)
  registro_paquete_pymes_pymes(request,seguimento)
  return &comms.Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *comms.Request_CrearOrdenRetail) (*comms.Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)
  seguimento:=registro_logico_retail("retail",request)
  registro_paquete_pymes_retail(request,seguimento)
  return &comms.Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, request *comms.Request_Seguimiento) (*comms.Response_Seguimiento, error) {
  log.Printf("Receive message %d", request.Seguimiento)

  for i:=0;i<6;i++{
    if(s.envios_s[i].Seguimiento==int(request.Seguimiento)){
      log.Printf("desde memoria :)")
      return &comms.Response_Seguimiento{Estado: s.envios_s[i].Estado}, nil
    }
  }
  aux:=strconv.Itoa(int(request.Seguimiento))
  csvFile,error:=os.Open("../storage/logica/"+aux+".csv")
  if error !=nil{
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
      return &comms.Response_Seguimiento{Estado: line[5]}, nil
    }
  }
  csvFile.Close()
  return &comms.Response_Seguimiento{Estado: "Esto no deberia suceder :)"}, nil
}

func (s *Server) SolicitarPaquete(ctx context.Context, request *comms.Request_SolicitarPaquete) (*comms.Response_SolicitarPaquete, error) {
  log.Printf("Receive message %s", request.Tipo)
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
  aux:=strconv.Itoa(int(x.Seguimiento))
  csvFile,_:=os.Open("../storage/logica/"+aux+".csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  line,_:=reader.Read()
  csvFile.Close()
  Updater("../storage/logica/"+aux+".csv","En camino")

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
  return &comms.Response_SolicitarPaquete{Id:x.Id,Seguimiento:int32(x.Seguimiento),Tipo:x.Tipo,Valor:int32(x.Valor),Tienda:line[5],Destino:line[6],}, nil
}

func Updater(n_file string,estado string){
  log.Printf("Seguimiento: %s", n_file)
  csvfile ,_:= os.Open(n_file)
  reader := csv.NewReader(bufio.NewReader(csvfile))
  line,_ :=reader.Read()
  csvfile.Close()
  change_id:=line[1]
  change_tipo:=line[2]
  nombrearch:="../storage/logica/retail.csv"
  switch change_tipo{
  case "retail":
    nombrearch="../storage/logica/retail.csv"
  default:
    nombrearch="../storage/logica/pymes.csv"
  }
  csvfile ,_= os.Open(nombrearch)
  reader = csv.NewReader(bufio.NewReader(csvfile))
  _=os.Remove("../storage/logica/aux.csv")
  csvfilex ,_:= os.OpenFile("../storage/logica/aux.csv", os.O_WRONLY|os.O_CREATE, 0777)
  writer:=csv.NewWriter(csvfilex)
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Fatal(error)
    }
    switch line[0] {
      case change_id:
        var guardar = [][]string{{line[0],line[1],line[2],line[3],line[4],estado},}
        error=writer.WriteAll(guardar)
      default:
        var guardar = [][]string{{line[0],line[1],line[2],line[3],line[4],line[5]},}
        error=writer.WriteAll(guardar)
    }
  }
  csvfilex.Close()
  csvfile.Close()
  Updater_csv("../storage/logica/aux.csv",nombrearch)
}

func Updater_csv(aux string, namefile string){
  csvfile ,_:= os.Open(aux)
  reader := csv.NewReader(bufio.NewReader(csvfile))
  csvfilex ,_:= os.OpenFile(namefile, os.O_WRONLY|os.O_CREATE, 0777)
  writer:=csv.NewWriter(csvfilex)
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
        log.Printf("gg po")
        log.Printf("texto %+v",line)
        log.Fatal(error)
        log.Printf("gg po 2")
    }
    var guardar = [][]string{{line[0],line[1],line[2],line[3],line[4],line[5]},}
    _=writer.WriteAll(guardar)
  }
  csvfilex.Close()
  csvfile.Close()
}

func LFP_R(pakete *paquete){
  file,_:=os.Open("../storage/logica/retail.csv")
  reader := csv.NewReader(bufio.NewReader(file))
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
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
  log.Printf("Receive message %s", request.Id)
  return &comms.Response_Estado{Recibido: "holo"}, nil
}

func (s *Server) LimpiarRegistros(ctx context.Context, request *comms.Dummy) (*comms.Dummy, error){
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
  return &comms.Dummy{Id:"1"}, nil

}

func main() {
  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := Server{}
  for i:=0;i<6;i++{
    s.envios_s[i].Uso="0"
  }
  grpcServer := grpc.NewServer()
  comms.RegisterCommsServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}
