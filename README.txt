Natalia        Herrera  Lora   201673501-1
Andrés Eduardo Shehadeh Gullón 201673560-7

-----------Como ejecutar-----------
En la carpeta del Makefile (por defecto ./repo), escribir en consola

"$Make [nombre]"

donde [nombre] corresponde al sistema que desea correr:
  -logistica
  -cliente
  -camión
  -finanzas
  --------------------------------------------


  -----------Cosas a considerar-----------
1.- Para mantener las ip's utilizadas en los códigos, se recomienda ejecutar los sistemas
    en las maquinas asignadas:
        logística --->maquina 93
        cliente   --->maquina 94
        camión    --->maquina 95
        finanzas  --->maquina 96

2.-Asegurese de tener corriendo el servidor de rabbit en la maquina 96

3.-Si usted quiere utilizar su propio input de pedidos de paquetes, remplace los archivos csv que se encuentran en la carpeta Pedidos ("pymes.csv" y/o "retail.csv")

4.-Para evitar una sobrecarga de información y que se repitan los Id de los pedidos, hicimos que al inicio de la ejecución del servidor se borren los Pedidos
   registrados

5.- al igual que el punto anterior, limpiamos el registro de finanzas y camiones, todo esto con la finalidad de que sea cómodo volver a ejecutar los sistemas

6.- Favor de no hacer más de 1 pedido con el mismo Id

7.-El sistema soporta la ejecución de 2 clientes, mientras estos sean uno de pyme y otro de retail, si usted quiere ejecutar otro cliente mientras
  el servidor está en ejecución tendrá que cambiar uno de los csv(pyme/retail), dado que si no lo hace el nuevo cliente hará pedidos con el mismo ID
  ya inscritos en el servidor.

8.- Al iniciar cliente, este pedirá 2 inputs, el tipo de cliente y tiempo entre request al servidor

9.- Al iniciar camión este pedirá los tiempos 2 de espera

10.-Todo funcionó correctamente, asique si hay algún problema no dude en contactarnos :D

11.-Si correrá los sistemas en otras máquinas haga los cambios necesarios en las conexiones, además de las instalaciones y los "go get [link]"
