const people = {
  _id: ObjectId("..."),
  name: "John Smith",
  email: "john.smith@example.com",
  phone: "+55 85 99999-9999",
  document: "123.456.789-00", // index
  birthDate: new Date("1980-05-15"),
};

const houseHolds = {
  _id: ObjectId("..."),
  persondocument: "123.456.789-00",
  meter: "MTR123456",
  address: {
    street: "Oak Avenue",
    number: 123,
    neighborhood: "Downtown",
    city: "Maracanaú",
    state: "CE",
  },
};

const readings = {
  _id: ObjectId("..."),
  meter: "MTR123456",
  readingDate: new Date("2025-06-01T00:00:00Z"),
  meterReading: 2453,
  monthlyConsumption: 300,
  estimatedCost: 270.0,
  recorded_by: ObjectId("..."),
};

const readers = {
  _id: ObjectId("..."),
  name: "Carlos Nunes",
  email: "carlos.nunes@readenergy.com",
  phone: "+55 85 98888-7777",
};


db.people.insertMany([
  {
    name: "Seiya de Pégaso",
    email: "saoriiiii@gmail.com",
    phone: "+55 85 99876-5432",
    document: "123.456.789-00",
    birthDate: new Date("1980-05-15")
  },
  {
    name: "Shiryu de Dragão",
    email: "greendragon9@gmail.com",
    phone: "+55 11 98888-8888",
    document: "987.654.321-00",
    birthDate: new Date("1990-08-22")
  },
  {
    name: "Hyoga de Cisne",
    email: "geladinho99@gmail.com",
    phone: "+55 21 97777-7777",
    document: "456.789.123-00",
    birthDate: new Date("1985-11-30")
  },
  {
    name: "Shun de Andrômeda",
    email: "andromede24@gmail.com",
    phone: "+55 31 96666-6666",
    document: "789.123.456-00",
    birthDate: new Date("1995-03-10")
  },
  {
    name: "Ikki de Fênix",
    email: "perfectfenix10@gmail.com",
    phone: "+55 51 95555-5555",
    document: "321.654.987-00",
    birthDate: new Date("1978-07-18")
  },
  {
    name: "Mu de Áries",
    email: "therealgoat80@gmail.com",
    phone: "+55 19 94444-4444",
    document: "654.321.987-00",
    birthDate: new Date("1992-12-05")
  },
  {
    name: "Aldebaran de Touro",
    email: "baraodogado777@gmail.com",
    phone: "+55 47 93333-3333",
    document: "234.567.890-00",
    birthDate: new Date("1988-09-25")
  },
  {
    name: "Saga de Gêmeos",
    email: "duality2@gmail.com",
    phone: "+55 27 92222-2222",
    document: "567.890.123-00",
    birthDate: new Date("1998-04-12")
  },
  {
    name: "Máscara da Morte de Câncer",
    email: "deathmask44@gmail.com",
    phone: "+55 81 91111-1111",
    document: "890.123.456-00",
    birthDate: new Date("1975-01-20")
  },
  {
    name: "Camus de Aquário",
    email: "ficafrioae00@gmail.com",
    phone: "+55 48 90000-0000",
    document: "345.678.901-00",
    birthDate: new Date("1983-06-08")
  }
])

db.readers.insertMany([
  {
    name: "Shaka de Virgem",
    email: "soleigeral777@gmail.com",
    phone: "+55 11 98888-1111",
    document: "121.131.141-11"
  },
  {
    name: "Afrodite de Peixes",
    email: "peixebolagato24@gmail.com",
    phone: "+55 21 98888-2222",
    document: "121.131.141-22"
  },
  {
    name: "Milo de Escorpião",
    email: "getoverhere@gmail.com",
    phone: "+55 19 98888-4444",
    document: "121.131.141-33"
  }
]);

db.houseHolds.insertMany([
  {
    persondocument: "123.456.789-00", 
    meter: "MTR789101",
    address: {
      street: "Rua Meteoro da Paixão",
      number: 88,
      neighborhood: "Lagoa dos Cavalo",
      city: "Beberibe",
      state: "CE"
    }
  },
  {
    persondocument: "987.654.321-00",
    meter: "MTR789102",
    address: {
      street: "Rua da Bica",
      number: 100,
      neighborhood: "Sitio Lucas",
      city: "Beberibe",
      state: "CE"
    }
  },
  {
    persondocument: "456.789.123-00",
    meter: "MTR789103",
    address: {
      street: "Rua 0",
      number: 23,
      neighborhood: "Cajueiro",
      city: "Aracati",
      state: "CE"
    }
  },
  {
    persondocument: "789.123.456-00",
    meter: "MTR789104",
    address: {
      street: "Beco mix",
      number: 9,
      neighborhood: "Centro",
      city: "Aracati",
      state: "CE"
    }
  },
  {
    persondocument: "321.654.987-00",
    meter: "MTR789105",
    address: {
      street: "Rua das Cinzas",
      number: 10,
      neighborhood: "Cacocity",
      city: "Beberibe",
      state: "CE"
    }
  },
  {
    persondocument: "654.321.987-00",
    meter: "MTR789106",
    address: {
      street: "Sitio do Bodão",
      number: 12,
      neighborhood: "Córrego do Retito",
      city: "Icapuí",
      state: "CE"
    }
  },
  {
    persondocument: "234.567.890-00", 
    meter: "MTR789107",
    address: {
      street: "Rua do Chifre",
      number: 7,
      neighborhood: "Loteamento",
      city: "Beberibe",
      state: "CE"
    }
  },
  {
    persondocument: "567.890.123-00", 
    meter: "MTR789108",
    address: {
      street: "Rua 2",
      number: 13,
      neighborhood: "Centro",
      city: "Aracati",
      state: "CE"
    }
  },
  {
    persondocument: "890.123.456-00", 
    meter: "MTR789109",
    address: {
      street: "Beco da Morte",
      number: 44,
      neighborhood: "Pedregal",
      city: "Aracati",
      state: "CE"
    }
  },
  {
    persondocument: "345.678.901-00", 
    meter: "MTR789110",
    address: {
      street: "rua das Profundezas",
      number: 65,
      neighborhood: "Pontal",
      city: "Fortim",
      state: "CE"
    }
  }
]);

db.readings.insertMany([
  {
    _id: ObjectId("78496ffb87460f483d69e401"),
    meter: "MTR789101",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 1500,
    monthlyConsumption: 200,
    estimatedCost: 180.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae") 
  },
  {
    _id: ObjectId("78496ffb87460f483d69e402"),
    meter: "MTR789101",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 1700,
    monthlyConsumption: 200,
    estimatedCost: 180.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf") 
  },
  {
    _id: ObjectId("78496ffb87460f483d69e403"),
    meter: "MTR789101",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 1900,
    monthlyConsumption: 200,
    estimatedCost: 180.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0") 
  },
  {
    _id: ObjectId("78496ffb87460f483d69e404"),
    meter: "MTR789102",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 2200,
    monthlyConsumption: 250,
    estimatedCost: 225.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e405"),
    meter: "MTR789102",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 2450,
    monthlyConsumption: 250,
    estimatedCost: 225.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e406"),
    meter: "MTR789102",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 2700,
    monthlyConsumption: 250,
    estimatedCost: 225.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e407"),
    meter: "MTR789103",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 1800,
    monthlyConsumption: 150,
    estimatedCost: 135.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e408"),
    meter: "MTR789103",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 1950,
    monthlyConsumption: 150,
    estimatedCost: 135.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e409"),
    meter: "MTR789103",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 2100,
    monthlyConsumption: 150,
    estimatedCost: 135.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40a"),
    meter: "MTR789104",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 2100,
    monthlyConsumption: 180,
    estimatedCost: 162.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40b"),
    meter: "MTR789104",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 2280,
    monthlyConsumption: 180,
    estimatedCost: 162.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40c"),
    meter: "MTR789104",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 2460,
    monthlyConsumption: 180,
    estimatedCost: 162.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40d"),
    meter: "MTR789105",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 3200,
    monthlyConsumption: 400,
    estimatedCost: 360.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40e"),
    meter: "MTR789105",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 3600,
    monthlyConsumption: 400,
    estimatedCost: 360.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e40f"),
    meter: "MTR789105",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 4000,
    monthlyConsumption: 400,
    estimatedCost: 360.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e410"),
    meter: "MTR789106",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 2800,
    monthlyConsumption: 220,
    estimatedCost: 198.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e411"),
    meter: "MTR789106",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 3020,
    monthlyConsumption: 220,
    estimatedCost: 198.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e412"),
    meter: "MTR789106",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 3240,
    monthlyConsumption: 220,
    estimatedCost: 198.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e413"),
    meter: "MTR789107",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 3800,
    monthlyConsumption: 350,
    estimatedCost: 315.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e414"),
    meter: "MTR789107",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 4150,
    monthlyConsumption: 350,
    estimatedCost: 315.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e415"),
    meter: "MTR789107",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 4500,
    monthlyConsumption: 350,
    estimatedCost: 315.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e416"),
    meter: "MTR789108",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 4200,
    monthlyConsumption: 380,
    estimatedCost: 342.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e417"),
    meter: "MTR789108",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 4580,
    monthlyConsumption: 380,
    estimatedCost: 342.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e418"),
    meter: "MTR789108",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 4960,
    monthlyConsumption: 380,
    estimatedCost: 342.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e419"),
    meter: "MTR789109",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 2900,
    monthlyConsumption: 240,
    estimatedCost: 216.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e41a"),
    meter: "MTR789109",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 3140,
    monthlyConsumption: 240,
    estimatedCost: 216.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e41b"),
    meter: "MTR789109",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 3380,
    monthlyConsumption: 240,
    estimatedCost: 216.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e41c"),
    meter: "MTR789110",
    readingDate: new Date("2025-01-01T00:00:00Z"),
    meterReading: 1950,
    monthlyConsumption: 120,
    estimatedCost: 108.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeae")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e41d"),
    meter: "MTR789110",
    readingDate: new Date("2025-02-01T00:00:00Z"),
    meterReading: 2070,
    monthlyConsumption: 120,
    estimatedCost: 108.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeaf")
  },
  {
    _id: ObjectId("78496ffb87460f483d69e41e"),
    meter: "MTR789110",
    readingDate: new Date("2025-03-01T00:00:00Z"),
    meterReading: 2190,
    monthlyConsumption: 120,
    estimatedCost: 108.0,
    recorded_by: ObjectId("684d7e4ebfdd00abe72cfeb0")
  }
]);

db.people.updateOne(
  { document: '123.456.789-12' },
  { $set: {
    email: "therealgoat01@gmail.com",
    phone: "+55 85 99876-2345"
  }}
)

db.people.findOne({document: "123.456.789-12"})

db.houseHolds.updateOne(
  { meter: "MTR7891777" },
  { $set: {
    address: {
      number: 1000
    }
  }}
)


