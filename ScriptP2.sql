CREATE TABLE temporada(
    id_temporada    NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,  
    nombre          VARCHAR(50),
    anyo            NUMBER (4),
    mes             NUMBER (2),
    fecha_fin       DATE,
    estado          CHAR(1)
);

CREATE TABLE fase (
    id_fase NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
    descripcion VARCHAR2(30)
);

CREATE TABLE jornada(
    id_jornada      NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,    
    nombre          VARCHAR(50),    
    fecha_inicio    DATE,
    fecha_fin       DATE,
    id_temporada    NUMBER,
    id_fase         NUMBER,
    FOREIGN KEY(id_fase) REFERENCES fase(id_fase),
    FOREIGN KEY(id_temporada) REFERENCES temporada(id_temporada)
);

CREATE TABLE deporte(
    id_deporte      NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,    
    nombre          VARCHAR(50),    
    imagen          VARCHAR(50),
    color           VARCHAR(50)
);

CREATE TABLE evento(
    id_evento       NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,  
    local         VARCHAR(50),
    visitante       VARCHAR(50),
    fecha           DATE  ,
    marcador_local  NUMBER,
    marcador_visita NUMBER,
    id_jornada      NUMBER,
    id_deporte      NUMBER,
    FOREIGN KEY(id_jornada) REFERENCES jornada(id_jornada),    
    FOREIGN KEY(id_deporte) REFERENCES deporte(id_deporte)
);

CREATE TABLE usuario(
    id_usuario      NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    nombre          VARCHAR(50),
    apellido        VARCHAR(50),
    clave           VARCHAR(50),
    usuario         VARCHAR(50),
    nacimiento      DATE,
    registro        DATE,
    email           VARCHAR(50),
    foto            VARCHAR(50)
);

CREATE TABLE membresia(
    id_membresia    NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,   
    nombre          VARCHAR(50), 
    precio          NUMBER
);

CREATE TABLE detalle_membresia(
    id_detalle      NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,  
    suscripcion     CHAR(1),
    id_usuario      NUMBER,
    id_membresia    NUMBER,  
    id_temporada    NUMBER,
    FOREIGN KEY(id_usuario) REFERENCES usuario(id_usuario),
    FOREIGN KEY(id_temporada) REFERENCES temporada(id_temporada),
    FOREIGN KEY(id_membresia) REFERENCES membresia(id_membresia)
);

CREATE TABLE prediccion(
    id_prediccion   NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    local         NUMBER,
    visitante       NUMBER,
    id_evento       NUMBER,
    id_usuario      NUMBER,
    FOREIGN KEY(id_evento) REFERENCES evento(id_evento),
    FOREIGN KEY(id_usuario) REFERENCES usuario(id_usuario)
);

--    alter SESSION set NLS_DATE_FORMAT = 'DD-MM-YYYY HH24:MI' 