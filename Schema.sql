USE CapitalOne;

CREATE TABLE Cliente (
    IdCliente INT PRIMARY KEY AUTO_INCREMENT,
    user VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE PropiedadesClientes (
    IdPropiedad INT PRIMARY KEY AUTO_INCREMENT,
    IdCliente INT,
    `Lat` DECIMAL(9,6),
    `Long` DECIMAL(9,6),
    Ciudad VARCHAR(100),
    Estado VARCHAR(100),
    Address VARCHAR(255),
    Precio DECIMAL(15,2),
    SqFt INT,
    TipoTerreno VARCHAR(50),
    Status VARCHAR(50),
    FOREIGN KEY (IdCliente) REFERENCES Cliente(IdCliente)
);

CREATE TABLE PropiedadesAsociadas (
    IdPropiedadAsociada INT PRIMARY KEY AUTO_INCREMENT,
    IdPropiedad INT,
    `Lat` DECIMAL(9,6),
    `Long` DECIMAL(9,6),
    Ciudad VARCHAR(100),
    Estado VARCHAR(100),
    Address VARCHAR(255),
    Precio DECIMAL(15,2),
    SqFt INT,
    TipoTerreno VARCHAR(50),
    Status VARCHAR(50),
    FOREIGN KEY (IdPropiedad) REFERENCES PropiedadesClientes(IdPropiedad)
);

