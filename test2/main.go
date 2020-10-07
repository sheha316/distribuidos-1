package main

import "fmt"

func main() {
  fmt.Println("Hello World!")

  elliot := &Person{
    Name: "Elliot",
    Age: 24,
    }

  data, err := proto.Marshal(elliot)
  if err != nil {
    log.Fatal("Marshalling error: ", err)
  }

  fmt.Println(data)

  newElliot := &Person{}
  err = proto.Unmarshal(data, newElliot)
  if err != nil {
    log.Fatal("Unmarshalling error: ", err)
  }

  fmt.Println(newElliot.getAge())
  fmt.Println(newElliot.getName())

}
