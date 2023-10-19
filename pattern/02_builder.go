package pattern

type DatabaseConnection struct {
    Host     string
    Port     int
    Username string
    Password string
    Database string
}

type DatabaseConnectionBuilder interface {
    SetHost(host string) DatabaseConnectionBuilder
    SetPort(port int) DatabaseConnectionBuilder
    SetUsername(username string) DatabaseConnectionBuilder
    SetPassword(password string) DatabaseConnectionBuilder
    SetDatabase(database string) DatabaseConnectionBuilder
    Build() *DatabaseConnection
}


type SomeDatabaseConnectionBuilder struct {
    connection *DatabaseConnection
}

func NewDatabaseConnectionBuilder() *SomeDatabaseConnectionBuilder {
    return &SomeDatabaseConnectionBuilder{connection: &DatabaseConnection{}}
}

func (b *SomeDatabaseConnectionBuilder) SetHost(host string) DatabaseConnectionBuilder {
    b.connection.Host = host
    return b
}

func (b *SomeDatabaseConnectionBuilder) SetPort(port int) DatabaseConnectionBuilder {
    b.connection.Port = port
    return b
}

func (b *SomeDatabaseConnectionBuilder) SetUsername(username string) DatabaseConnectionBuilder {
    b.connection.Username = username
    return b
}

func (b *SomeDatabaseConnectionBuilder) SetPassword(password string) DatabaseConnectionBuilder {
    b.connection.Password = password
    return b
}

func (b *SomeDatabaseConnectionBuilder) SetDatabase(database string) DatabaseConnectionBuilder {
    b.connection.Database = database
    return b
}

func (b *SomeDatabaseConnectionBuilder) Build() *DatabaseConnection {
    return b.connection
}

//Строитель — это порождающий паттерн проектирования, который позволяет создавать объекты пошагово.
//Плюсами являются упрощение создания объектов и гибкость изменения их конфигураций
//Минусы - увеличение и усложнение кода
