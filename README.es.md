Proyecto Domlan
===============

Domlan es un servidor para Home Atomation/Domótica multi-platforma.
El protocolo principal soportado por ahora es MQTT

## Compilar y correr

Se debe tener instalado Go 1.4.x, Git, Mercurial, Bower (módulo de Node.js) y un broker MQTT (por ejemplo [Mosquitto](http://mosquitto.org/))

```bash
~$ git clone https://github.com/ljgago/domlan.git && cd domlan
~$ ./build.sh
~$ ./domlan
```

