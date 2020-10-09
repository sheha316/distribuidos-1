package comms
import (
  "log"
  "golang.org/x/net/context"
  "os"
  "io"
  "encoding/csv"
  "fmt"
  "strconv"
  "bufio"
)

type Server struct {
}

type Pedido_pymes_l struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
  Prioritario int
  Estado string
}

type Pedido_retail_l struct{
  Id string
  Producto string
  Valor int
  Tienda string
  Destino string
  Estado string
}
func (s *Server) CrearOrdenPyme(ctx context.Context, request *Request_CrearOrdenPyme) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)

  seguimento:=0
  file,erros:=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  }
  file,erros=os.Create("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  if erros!=nil{
    log.Printf("error:")
    fmt.Println(erros)
    log.Printf("fin:")
  }
  writer:=csv.NewWriter(file)
  var guardar = [][]string{
    {request.Id,request.Producto,string(request.Valor),request.Tienda,request.Destino,string(request.Prioritario),"En bodega"},
  }
  erros=writer.WriteAll(guardar)
  aux:="1"+strconv.Itoa(seguimento)
  seguimento,_=strconv.Atoi(aux)
  file.Close()
  return &Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) CrearOrdenRetail(ctx context.Context, request *Request_CrearOrdenRetail) (*Response_CrearOrden, error) {
  log.Printf("Receive message %s", request.Id)

  seguimento:=0
  file,erros:=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    seguimento++
    file.Close()
    file,erros=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  }
  file,erros=os.Create("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  if erros!=nil{
    log.Printf("error:")
    fmt.Println(erros)
    log.Printf("fin:")
  }
  writer:=csv.NewWriter(file)
  var guardar = [][]string{
    {request.Id,request.Producto,string(request.Valor),request.Tienda,request.Destino,"En bodega"},
  }
  erros=writer.WriteAll(guardar)
  aux:="2"+strconv.Itoa(seguimento)
  seguimento,_=strconv.Atoi(aux)
  file.Close()
  return &Response_CrearOrden{Seguimiento: int32(seguimento)}, nil
}

func (s *Server) Seguimiento(ctx context.Context, request *Request_Seguimiento) (*Response_Seguimiento, error) {
  log.Printf("Receive message %d", request.Seguimiento)
  aux:=strconv.Itoa(int(request.Seguimiento))
  csvFile,error:=os.Open("./paquetes/"+aux+".csv")
  if error !=nil{
    return &Response_Seguimiento{Estado: "Paquete no existe"}, nil
  }
  reader := csv.NewReader(bufio.NewReader(csvFile))
  for{
    line,error :=reader.Read()
    if error==io.EOF{
      break
    }else if error!=nil{
      log.Fatal(error)
    }
    aux1,_:=strconv.Atoi(line[2])
    switch  string(aux[0]){
      case "1":
        var pedido []Pedido_pymes_l
        aux2,_:=strconv.Atoi(line[5])
        pedido=append(pedido,Pedido_pymes_l{
          Id:line[0],
          Producto:line[1],
          Valor:aux1,
          Tienda:line[3],
          Destino:line[4],
          Prioritario:aux2,
          Estado:line[6],
        })
        return &Response_Seguimiento{Estado: pedido[0].Estado}, nil
      default:
        var pedido []Pedido_retail_l
        pedido=append(pedido,Pedido_retail_l{
          Id:line[0],
          Producto:line[1],
          Valor:aux1,
          Tienda:line[3],
          Destino:line[4],
          Estado:line[5],
        })
        return &Response_Seguimiento{Estado: pedido[0].Estado}, nil

    }
  }
  return &Response_Seguimiento{Estado: "Esto no deberia suceder :)"}, nil
}

func (s *Server) SolicitarPaquete(ctx context.Context, request *Request_SolicitarPaquete) (*Response_SolicitarPaquete, error) {
  log.Printf("Receive message %s", request.Tipo)
  x:=&Pedido_retail_l{Valor: -1,}
  var tipo_p string
  //retail=2 pyme=1
  switch request.Tipo {
  case "retail":
    LFP_R(x)
    tipo_p="retail"
    if(x.Valor==-1){
      LFP_P(x)
      tipo_p="prioritario"
    }

  default:
    LFP_P(x)
    tipo_p="prioritario"
    if(int(x.Valor)==-1){
      LFP_N(x)
      tipo_p="normal"
    }
  }
  return &Response_SolicitarPaquete{Id:x.Id,Tipo:tipo_p,Valor: int32(x.Valor),Tienda: x.Tienda,Destino: x.Destino}, nil
}
func Updater(n_file string,estado string,tipo string){
  csvfile ,_:= os.Open(n_file)
  reader := csv.NewReader(bufio.NewReader(csvfile))
  line,error :=reader.Read()
  csvfile.Close()
  csvfilex ,_:= os.OpenFile(n_file, os.O_WRONLY|os.O_CREATE, 0777)
  writer:=csv.NewWriter(csvfilex)
  switch tipo{
  case "retail":
    var guardar =[][]string{{line[0],line[1],line[2],line[3],line[4],estado,},}
    erros:=writer.WriteAll(guardar)
  default:
    var guardar =[][]string{{line[0],line[1],line[2],line[3],line[4],line[5],estado,},}
    erros:=writer.WriteAll(guardar)
  }
  csvfilex.Close()
}

func LFP_R(pakete *Pedido_retail_l){
  seguimento:=0
  file,erros:=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    reader := csv.NewReader(bufio.NewReader(file))
    line,error :=reader.Read()
    file.Close()
    if(line[5]=="En bodega"){
      aux1,_:=strconv.Atoi(line[2])
      pakete.Id=line[0]
      pakete.Valor=aux1
      pakete.Tienda=line[3]
      pakete.Destino=line[4]
      Updater("./paquetes/2"+strconv.Itoa(seguimento)+".csv","En camino","retail")
      return
    }
    seguimento++
    file,erros=os.Open("./paquetes/2"+strconv.Itoa(seguimento)+".csv")
  }
  pakete.Valor=-1
}

func LFP_P(pakete *Pedido_retail_l){
  seguimento:=0
  file,erros:=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    reader := csv.NewReader(bufio.NewReader(file))
    line,error :=reader.Read()
    file.Close()
    if(line[6]=="En bodega" && line[5]=="1"){
      aux1,_:=strconv.Atoi(line[2])
      pakete.Id=line[0]
      pakete.Valor=aux1
      pakete.Tienda=line[3]
      pakete.Destino=line[4]
      Updater("./paquetes/1"+strconv.Itoa(seguimento)+".csv","En camino","prioritario")
      return
    }
    seguimento++
    file,erros=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  }
  pakete.Valor=-1
}

func LFP_N(pakete *Pedido_retail_l){
  seguimento:=0
  file,erros:=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  for erros==nil{
    reader := csv.NewReader(bufio.NewReader(file))
    line,error :=reader.Read()
    file.Close()
    if(linea[6]=="En bodega" && linea[5]=="0"){
      aux1,_:=strconv.Atoi(line[2])
      pakete.Id=line[0]
      pakete.Valor=aux1
      pakete.Tienda=line[3]
      pakete.Destino=line[4]
      Updater("./paquetes/1"+strconv.Itoa(seguimento)+".csv","En camino","normal")
      return
    }
    seguimento++
    file,erros=os.Open("./paquetes/1"+strconv.Itoa(seguimento)+".csv")
  }
  pakete.Valor=-1
}





func (s *Server) InformarEstado(ctx context.Context, request *Request_Estado) (*Response_Estado, error) {
  log.Printf("Receive message %s", request.Id)
  return &Response_Estado{Recibido: "holo"}, nil
}
