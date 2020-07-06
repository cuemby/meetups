# episodio 8 - Introducción a Ingress Controller con NGINX

## que es un ingress

Nos permite manejar el trafico externo hacia los servicios del cluster,
expone rutas fuera de nuestro cluster y las direcciona hacia los servicios que tenemos defino
a traves de los recursos `Ingress`.

## diferencia/ventajas entre un ingress y loadBalancer

---- LoadBalancer
* Nos permite exponer un servicio a internet, mediante una IP que dirije el trafico al servicio y este 
lo reparte entre los pods disponibles

* Si usamos microservicios, con este metodo deberiamos crear un loadBalancer para exponer cada uno de estos.

* soporta practicamente todos los protocolos (HTTPs, gRPC, RCP)

--- Ingress
* El ingress no es un tipo de servicio, se ubica por delante de los servicios y funciona como un "enrutador".
Su capacidad depende del controlador instalador, por lo general cada proveedor trae uno propio por defecto.

* Soportan principalmente HTTP(S). 

* El controlador crea una service de LoadBalancer por el cual enrutaremos todo el trafico a donde corresponda.

## alternivas al ingress nginx controller 
- tabla comparativa entre las alternativas

## instalacion NGINX Ingress controller

website https://kubernetes.github.io/ingress-nginx/
```bash
# for digital ocean
kubectl create ns ingress-nginx
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/do/deploy.yaml
```

## explicacion del ejemplo
(hacer un diagrama de como va a funcionar la app)

www = page 
app = dashboard
api = apis  (accounts, todo)

## desplegamos los servicios

```bash
kubectl apply -f servicios/
```

## configuramos los ingress
```bash
kubectl apply -f ingress/
```

## instalamos los secrets para los SSL
```bash
kubectl apply -f ssl/
```
## configuramos los ingress con SSL
```bash
kubectl apply -f ingress-ssl/
```

## create tls secret 
kubectl create secret tls CERT_NAME --key KEY_FILE --cert CERT_FILE