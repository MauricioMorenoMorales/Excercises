La estructura de las bases de datos es la siguiente

Software Base de datos > Base de datos > Esquemas > Tablas > Data

Un software de base de datos puede tener varias bases de datos
Un esquema es un grupo de tablas dentro de una base de datos, tambien puede contener configuraciones o metadatos, y sirven para organizar la informacion
En MySQL Database y Schema significan lo mismo pero en otros sistemas como PostgreSQL son cosas diferentes
Las tablas son lo más importante dentro de una base de datos:
Cada tabla almacena datos de un tipo específico, como usuarios, productos, órdenes, etc.
Dentro de cada tabla se definen:
las columnas (nombre, tipo de dato, tamaño, valores por defecto, etc.)
y cada fila representa una entrada individual de

## Crear un squema
```sql
CREATE SCHEMA `new_schema` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

