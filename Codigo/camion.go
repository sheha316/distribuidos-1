package main
import (
  "log"
  "github.com/sheha316/distribuidos-1/Codigo/comms"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "time"
  "strconv"
  "math/rand"
  "fmt"
  "encoding/csv"
  "os"
)
type paquete_info struct{
  Id string
  Tipo string
  Valor int
  Tienda string
  Destino string
  Intentos int
  Fecha string
}
type Camion struct{
  Tipo string
  Id string
  Paquete_inf1 paquete_info
  Paquete_inf2 paquete_info
  Paquetes int
  Estado int
}

func request_paquete_1(conn *grpc.ClientConn, kamion *Camion){
  c := comms.NewCommsClient(conn)
  response, _ := c.SolicitarPaquete(context.Background(), &comms.Request_SolicitarPaquete{Tipo: kamion.Tipo})
  if(int(response.Valor)!=-1){
    kamion.Paquete_inf1=paquete_info{Id:response.Id,Tipo:response.Tipo,Valor:int(response.Valor),
    Tienda:response.Tienda,Destino:response.Destino,Intentos:0,Fecha:"0"}
    kamion.Paquetes=1
  }
}
func request_paquete_2(conn *grpc.ClientConn, kamion *Camion){
  c := comms.NewCommsClient(conn)
  response, _ := c.SolicitarPaquete(context.Background(), &comms.Request_SolicitarPaquete{Tipo: kamion.Tipo,Id:kamion.Id})
  if(int(response.Valor)!=-1){
    kamion.Paquete_inf2=paquete_info{Id:response.Id,Tipo:response.Tipo,Valor:int(response.Valor),
    Tienda:response.Tienda,Destino:response.Destino,Intentos:0,Fecha:"0"}
    kamion.Paquetes=2
  }
}

func cargar_camion(conn *grpc.ClientConn, kamion *Camion,tiempo int){
  request_paquete_1(conn , kamion)
  if(kamion.Paquetes==1){
    time.Sleep(time.Duration(tiempo) * time.Second)
    request_paquete_2(conn , kamion)
    kamion.Estado=1
  }
}

func expirado(paquete paquete_info)(bool){
  if(paquete.Fecha!="0"){
    return true
  }
  if( (paquete.Tipo=="retail") && (paquete.Intentos<3)){
    return false
  }
  if( (paquete.Tipo=="normal"||paquete.Tipo=="prioritario") && (paquete.Intentos<2) && (paquete.Intentos*10<=paquete.Valor)){
    return false
  }
  return true
}

func Reparto(kamion *Camion,tiempo int){
  kamion.Estado=0
  if(kamion.Paquetes==2){
    if(kamion.Paquete_inf1.Valor>kamion.Paquete_inf2.Valor){
      if(!expirado(kamion.Paquete_inf1)){
        time.Sleep(time.Duration(tiempo) * time.Second)
        kamion.Estado=1
        kamion.Paquete_inf1.Intentos++
        if(rand.Intn(100)<80){
          kamion.Paquete_inf1.Fecha=time.Now().String()
        }
      }
      if(!expirado(kamion.Paquete_inf2)){
        time.Sleep(time.Duration(tiempo) * time.Second)
        kamion.Paquete_inf2.Intentos++
        kamion.Estado=1
        if(rand.Intn(100)<80){
          kamion.Paquete_inf2.Fecha=time.Now().String()
        }
      }
    }else if(kamion.Paquete_inf2.Valor>=kamion.Paquete_inf1.Valor){
      if(!expirado(kamion.Paquete_inf2)){
        time.Sleep(time.Duration(tiempo) * time.Second)
        kamion.Paquete_inf2.Intentos++
        kamion.Estado=1
        if(rand.Intn(100)<80){
          kamion.Paquete_inf2.Fecha=time.Now().String()
        }
      }
      if(!expirado(kamion.Paquete_inf1)){
        time.Sleep(time.Duration(tiempo) * time.Second)
        kamion.Paquete_inf1.Intentos++
        kamion.Estado=1
        if(rand.Intn(100)<80){
          kamion.Paquete_inf1.Fecha=time.Now().String()
        }
      }
    }
  }else if(kamion.Paquetes==1){
    if(!expirado(kamion.Paquete_inf1)){
      time.Sleep(time.Duration(tiempo) * time.Second)
      kamion.Paquete_inf1.Intentos++
      kamion.Estado=1
      if(rand.Intn(100)<80){
        kamion.Paquete_inf1.Fecha=time.Now().String()
      }
    }
  }
}

func registrar_paquete(id string,paquete paquete_info){
  f, err := os.OpenFile("./storage/camion/"+id+".csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    log.Printf("registrar_paquete")
    log.Fatal(err)
  }
  defer f.Close()
  var data [][]string
  data = append(data, []string{paquete.Id,paquete.Tipo,strconv.Itoa(paquete.Valor),paquete.Tienda,paquete.Destino,strconv.Itoa(paquete.Intentos),paquete.Fecha})
  w := csv.NewWriter(f)
  w.WriteAll(data)
  if err := w.Error(); err != nil {
    log.Printf("registrar_paquete")
    log.Fatal(err)
  }
}

func reporte(conn *grpc.ClientConn,kamion *Camion){
  c := comms.NewCommsClient(conn)
  var estadorm string
  for i:=0;i<kamion.Paquetes;i++{
    estadorm="Recibido"
    switch i {
      case 0:
        if(kamion.Paquete_inf1.Fecha=="0"){
          estadorm="No Recibido"
        }
        registrar_paquete(kamion.Id,kamion.Paquete_inf1)
        _, _ = c.InformarEstado(context.Background(), &comms.Request_Estado{Id:kamion.Paquete_inf1.Id,
                                                                                    Intentos:int32(kamion.Paquete_inf1.Intentos),
                                                                                    Fecha:kamion.Paquete_inf1.Fecha,
                                                                                    Estado:estadorm})
      case 1:
        if(kamion.Paquete_inf2.Fecha=="0"){
          estadorm="No Recibido"
        }
        registrar_paquete(kamion.Id,kamion.Paquete_inf2)
        _, _ = c.InformarEstado(context.Background(), &comms.Request_Estado{Id:kamion.Paquete_inf2.Id,
                                                                                    Intentos:int32(kamion.Paquete_inf2.Intentos),
                                                                                    Fecha:kamion.Paquete_inf2.Fecha,
                                                                                    Estado:estadorm})
    }
  }
  kamion.Paquetes=0
}

func superprint_fs(kamion *Camion){
  log.Printf("Camion: "+kamion.Id+",Tipo:"+kamion.Tipo+",Paquetes:"+strconv.Itoa(kamion.Paquetes))
  if(kamion.Paquetes==1){
    log.Printf("Paquete 1: %+v",kamion.Paquete_inf1)
  }else if(kamion.Paquetes==2){
    log.Printf("Paquete 1: %+v",kamion.Paquete_inf1)
    log.Printf("Paquete 2: %+v",kamion.Paquete_inf2)
  }
}
func superprint_ts(kamion *Camion){
  log.Printf("Camion: "+kamion.Id+",Tipo:"+kamion.Tipo+",Paquetes:"+strconv.Itoa(kamion.Paquetes))
  if(kamion.Paquetes==1){
    log.Printf("Paquete 1: %+v",kamion.Paquete_inf1)
  }else if(kamion.Paquetes==2){
    log.Printf("Paquete 1: %+v",kamion.Paquete_inf1)
    log.Printf("Paquete 2: %+v",kamion.Paquete_inf2)
  }
}

func main() {
  os.Remove("./storage/camion/1.csv")
  os.Remove("./storage/camion/2.csv")
  os.Remove("./storage/camion/3.csv")
  os.Create("./storage/camion/3.csv")
  os.Create("./storage/camion/2.csv")
  os.Create("./storage/camion/1.csv")
  camion_1:=&Camion{
    Tipo: "retail",Paquetes:0,Estado:0,Id:"1"}
  camion_2:=&Camion{
    Tipo: "retail",Paquetes:0,Estado:0,Id:"2"}
  camion_3:=&Camion{
    Tipo: "normal",Paquetes:0,Estado:0,Id:"3"}
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist93:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  var T_sp int
  var T_ep int
  log.Printf("Tiempo de espera de 2do paquete")
  fmt.Scanln(&T_sp)
  log.Printf("Tiempo que demora en entregar un paquete")
  fmt.Scanln(&T_ep)
  for{
      if(camion_1.Estado==0){
        cargar_camion(conn,camion_1,T_sp)
      }
      if(camion_2.Estado==0){
        cargar_camion(conn,camion_2,T_sp)
      }
      if(camion_3.Estado==0){
        cargar_camion(conn,camion_3,T_sp)
      }
      log.Printf("----Paquetes Recibidos----")
      superprint_fs(camion_1)
      superprint_fs(camion_2)
      superprint_fs(camion_3)
      log.Printf("--------------------------")/**/
      for camion_1.Estado==1{
        Reparto(camion_1,T_ep)
      }
      for camion_2.Estado==1{
        Reparto(camion_2,T_ep)
      }
      for camion_3.Estado==1{
        Reparto(camion_3,T_ep)
      }
      log.Printf("----Paquetes Entregados----")
      superprint_ts(camion_1)
      superprint_ts(camion_2)
      superprint_ts(camion_3)
      log.Printf("--------------------------")/**/
      if(camion_1.Paquetes!=0){
        reporte(conn,camion_1)
      }
      if(camion_2.Paquetes!=0){
        reporte(conn,camion_2)
      }
      if(camion_3.Paquetes!=0){
        reporte(conn,camion_3)
      }
    }
}
