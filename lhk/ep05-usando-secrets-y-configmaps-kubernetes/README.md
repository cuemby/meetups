# Episodio 05 - Usando secrets y configmaps en Kubernetes

Para lograr implementaciones continuas más rápidas, Microservices debe evitar la codificación rígida de los parámetros específicos de la aplicación en el código. La práctica recomendada por la aplicación de 12 factores es separar los parámetros de configuración del código de la aplicación. El manejo de configuraciones de aplicaciones como la ruta de almacenamiento, las cadenas de conexión db, las URL, las funciones de registro, en entornos de preparación de CICD puede ser doloroso, si no se administra adecuadamente.

Aunque el archivo Docker admite la configuración a través de variables de entorno, no se recomienda recrear imágenes para actualizaciones menores. Kubernetes tiene la respuesta: 'Mapa de configuración' con el que se pueden configurar los parámetros de la aplicación durante las implementaciones.

Un mapa de configuración es un diccionario de pares clave-válvula que se puede inyectar en las cápsulas de Kubernetes. Para saber más te esperamos en nuestro próximo Live Streaming.

## Agenda

* Introducción a configmap
* Como crear un configmap usando la linea de comandos y archivos `.yaml`
* Como crear un secret usando la linea de comandos y archivos `.yaml`
* Mapeando configmaps y secrets como variables de entorno en un Pod.
* Montando configmaps y secrets como directorios o archivos en un Pod.
* Algunos tips de seguridad para secrets!
* Preguntas y respuestas

## Comandos

```shell
kubectl create configmap my-config1 --from-file=configmaps/files
kubectl create configmap my-config2 --from-file=cfg1=configmaps/files/cfg1.txt --from-file=cfg2=configmaps/files/cfg2.txt
kubectl create configmap my-config3 --from-file=configmaps/files/cfg3.text
kubectl create configmap my-config4 --from-literal=saludo=hola --from-literal=despedida=chao
kubectl create configmap my-config5 --from-env-file=local.env
kubectl apply -f configmap.yaml
```

```shell
kubectl create secret generic my-secret1 --from-file=secrets/files
kubectl create secret generic my-secret2 --from-file=ssh-privatekey=files/id_rsa --from-file=ssh-publickey=files/id_rsa.pub
kubectl create secret generic my-secret3 --from-literal=user=juanito --from-literal=passwd=sshhh
kubectl create secret generic my-secret4 --from-file=ssh-privatekey=secrets/files/id_rsa --from-literal=passphrase=sshhh
kubectl create secret generic my-secret5 --from-env-file=secrets/local.env
kubectl apply -f secret.yaml
```
