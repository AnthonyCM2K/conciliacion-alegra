<!-- # Logica

## Ejemplo 1

1. Crear un mapa con los datos `id, datetime, Amount, email` de una factura.
2. La idea es que la variable `invoice` sea un mapa vacío y se llene con los datos obtenidos de la respuesta de la API.
3. Almacenamos todos las facturas en un slice.

## Ejemplo 2

1. Crear un struct con los datos `id, datetime, Amount, email` de una factura.
2. La idea es que la variable invoice sea un struct vacío y se llene con los datos obtenidos de la respuesta de la API.
3. Almacenamos todos las facturas en un slice.

## Pasos a seguir

1. Consumir las APIs de Alegra y el CMS.
   - La respuesta de las APIs está en formato JSON.
2. Deserializar (convertir ese JSON a una estructura) el JSON.
   - Crear las estructuras que representen los datos del JSON.
3. Extraer los campos `id, datetime, Amount, email` de las estructuras definidas que contienen los datos de Alegra y CMS.
4. Almacenar en un slice la estructura o el mapa con los datos extraídos.
   - Este paso debe hacerse tanto con los datos de Alegra como con los del CMS.
5. Una vez tengamos los dos slices con los datos extraídos de Alegra y del CMS, procedemos a recorrerlos y compararlos.

```go

```

## Validaciones que hicimos

1. Un slice a modo de reporte que contiene `notInCMS`: Facturas de Alegra que no estan en el CMS
2. Un slice a modo de reporte que contiene `notInAlegra`: Facturas de CMS que no estan en el Alegra
   - Se compara que los ids sean iguales, luego se hace la comparación de si los totales son diferentes; Con el objetivo de poder capturar las facturas en el reporte y tener el Amount.

```shell
		
```



