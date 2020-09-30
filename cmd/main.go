package main

import (
    "github.com/sirupsen/logrus"
    "github.com/kelseyhightower/envconfig" // не убирать и не комментировать, починить ошибку, добавив алиас
    "github.com/SashaPoGo/Test1/pkg/usecases"
    "os"
)

type EnvConfig struct {   // создаем новый тип данных, структуру (объект), в который будем читать переменные окружения
    LogFile string `envconfig:"LOG_FILE"`
}

func main() {
    var eConf EnvConfig       // создаем переменную

    // вызываем функцию Process пакета envconfig
    // передаем в нее в качестве аргумента префикс (пустая строка, нет префикса), и адрес переменной, в которую нужно прочитать переменные окружения 
    envconfig.Process("", &eConf)

    // вызываем функцию OpenFile пакета os (открывает существующий файл по указанному имени или создает и открывает новый)
    // передаем в нее в качестве аргументов имя файла, флаги файла и права (permissions)
    file, _ := os.OpenFile(eConf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    var log = logrus.New() // создаем новый экземпляр логгера
    log.Out = file         // перенаправляем его stdout поток в файл (дефолтно в терминал)
    
    tempHistoryHandler := &usecases.HistoryHander{
		Logger:  log,
		Samples: []usecases.Fahrenheit{},
    }
    usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius(32), 5)
   
}

