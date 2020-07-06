# Episodio 4 - Un vistazo a Prometheus

## link al video

[https://www.youtube.com/watch?v=bwh1urfG56Y](https://www.youtube.com/watch?v=bwh1urfG56Y)
## introducción

Prometheus es un sistema de monitoreo y alertas de codigo abierto 
(licencia Apache 2.0) desarrollado principalmente en Go por la empresa soundCloud. 
Fue lanzado en 2012 y muchas compañias y orgnizaciones lo han adoptado.

Prometheus se vinculo a la Cloud Native Computing Foundation (CNCF) en 2016. 

Prometheus no es un logger, su funcion es coleccionar datos cuantitativos. Ejemplo, la causa de una falla de un bloque de codigo no seria un dato para prometheus, pero el contar cuantas veces ha fallado, si es un dato que deberiamos enviar a prometheus.

Si lo que desea es hacer un logger, te recomiendo usar elasticsearch y kibana.




### como almacena los datos?

Prometheus almacena todos los datos como series de tiempo, es decir valores con fecha/hora asociadas a aun conjunto de etiquetas.
Cada serie de tiempo se identifica de manera unica por su nombre de metrica (ejempl o: `http_requests_total`) y pares opcionales
de clave/valor llamados etiquetas.

Prometheus crea un archivo local por serie de tiempo almacenando datos a medida que ocurren los scrape o se evaluan reglas.


### ¿En qué idioma está escrito Prometeo?
La mayoría de los componentes de Prometheus están escritos en Go. Algunos también están escritos en Java, Python y Ruby.

## Arquitectura

[Arquitectura](img/arquitecture.png)

## instalación

usaremos [helm](https://helm.sh/docs/intro/install/) para instalar prometheus

```bash
$ kubectl create ns monitoring
$ helm template prometheus stable/prometheus --namespace monitoring
```
## consultar metricas

Prometheus proporciona un lenguaje de consulta funcional llamado PromQL (Prometheus Query Language) que permite al usuario seleccionar y agregar datos de series de tiempo en tiempo real. El resultado de una expresión puede mostrarse como un gráfico, verse como datos tabulares en el navegador de expresiones de Prometheus o ser consumido por sistemas externos a través de la API HTTP. Ejemplos de consultas:

```
http_requests_total{}
rate(http_requests_total[5m])[30m:1m]
```

## ver dashboard de prometheus
```bash
# buscar el pod del server y hacer port-forward
 kubectl -n monitoring port-forward svc/prometheus-server 80
```

## graficar metricas

Para graficar metricas prometheus nos recomienda usar grafana.

```
helm install grafana stable/grafana
```
Para visualizar las metricas usaremos el siguiente dashboard:

```
Kubernetes cluster overview - id: 11802
Kubernetes cluster monitoring (via Prometheus) : 315
```

## desplegar aplicacion de ejemplo

https://github.com/brancz/prometheus-example-app




## Que puedo integrar a prometheus?

Hay un gran numero de de librerias y servidores que pueden ayudar a exporar informacion a prometheus, como por ejemplo metricas de linux, motores de bases de datos, etc. 

https://prometheus.io/docs/instrumenting/exporters/


## comparaciones con otros proyectos

- Graphite : Seria una mejor opcion si desea contener datos historicos a largo plazo.

- InfluxDB: Es uan base de datos de series de tiempo creada para moninotear metricas y eventos. Seria una buena opcion si se esta requiere un registro de eventos

- OpenTSDB: Es muy similar a prometheus