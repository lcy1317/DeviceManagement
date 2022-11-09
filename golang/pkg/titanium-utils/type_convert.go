package titaniumutils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// OneNumberto0Number 用于将一个个位数字前面加个0
func OneNumberto0Number(k int) string {
	if k < 10 {
		return "0" + strconv.Itoa(k)
	}
	return strconv.Itoa(k)
}

const (
	// January 1
	January time.Month = 1 + iota
	// February 2
	February
	// March 3
	March
	// April 4
	April
	// May 5
	May
	// June 6
	June
	// July 7
	July
	// August 8
	August
	// September 9
	September
	// October 10
	October
	// November 11
	November
	// December 12
	December
)

// GetNowTimeinString get now string
func GetNowTimeinString() string {
	t1 := OneNumberto0Number(time.Now().Year()) //年
	m := time.Now().Month()
	timeMonth := 1
	switch m.String() {
	case "January":
		timeMonth = 1
	case "February":
		timeMonth = 2
	case "March":
		timeMonth = 3
	case "April":
		timeMonth = 4
	case "May":
		timeMonth = 5
	case "June":
		timeMonth = 6
	case "July":
		timeMonth = 7
	case "August":
		timeMonth = 8
	case "September":
		timeMonth = 9
	case "October":
		timeMonth = 10
	case "November":
		timeMonth = 11
	case "December":
		timeMonth = 12
	default:
		timeMonth = 1
	}
	t2 := OneNumberto0Number(timeMonth)           //月
	t3 := OneNumberto0Number(time.Now().Day())    //日
	t4 := OneNumberto0Number(time.Now().Hour())   //小时
	t5 := OneNumberto0Number(time.Now().Minute()) //分钟
	t6 := OneNumberto0Number(time.Now().Second()) //秒

	return t1 + t2 + t3 + t4 + t5 + t6
}

// GetNowTimeNanoSecinString get now nano string
func GetNowTimeNanoSecinString() string {
	t1 := OneNumberto0Number(time.Now().Year())
	t2 := time.Now().Month().String()
	t3 := OneNumberto0Number(time.Now().Day())    //日
	t4 := OneNumberto0Number(time.Now().Hour())   //小时
	t5 := OneNumberto0Number(time.Now().Minute()) //分钟
	t6 := OneNumberto0Number(time.Now().Second()) //秒
	t7 := OneNumberto0Number(time.Now().Nanosecond())

	return t1 + t2 + t3 + t4 + t5 + t6 + t7
}

// Hexstring2Decint64 改函数用于将0x或者不是0x开头的十六进制数转成十进制的int64
func Hexstring2Decint64(val string) int64 {
	// 去掉0x / 0X

	if len(val) <= 2 {
		return 0
	}
	if val[0:2] == "0x" || val[0:2] == "0X" {
		val = val[2:]
	}
	fmt.Println(val)
	n, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		fmt.Println(err)
		logrus.Info("Hex convert to Dec Error!")
	}
	return n
}

// WriteLogNumber write log number
func WriteLogNumber(num int) {
	filePath := "./taiduutils/lognumber.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("check log number error %v \n", err)
		return
	}
	//及时关闭
	defer file.Close()
	//写入内容
	writer := bufio.NewWriter(file)
	writer.WriteString(strconv.Itoa(num))
	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}

// GetLogNumber get log number
func GetLogNumber() int {
	strLogNumber, err := ioutil.ReadFile("./taiduutils/lognumber.txt")
	logNumber, err := strconv.Atoi(string(strLogNumber))
	if err != nil {
		logrus.Info("Get Log Number Failed!")
	}
	fmt.Println(logNumber)
	return logNumber
}

// BytesToInt bytes to int
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

// BigIntToUint64 bitint to uint64
func BigIntToUint64(bigNum *big.Int) (uint64, error) {
	boolean := bigNum.IsUint64()
	if !boolean {
		return 0, fmt.Errorf("bigNum %v can't transfer to Uint64", bigNum)
	}
	return bigNum.Uint64(), nil
}

// BigIntToInt64 bitint to int64
func BigIntToInt64(bigNum *big.Int) (int64, error) {
	boolean := bigNum.IsInt64()
	if !boolean {
		return 0, fmt.Errorf("BigIntToint64 failed, bigNum %v can't transfer to Int64", bigNum)
	}
	return bigNum.Int64(), nil
}

// Uint64ToInt64 to int64
func Uint64ToInt64(num uint64) (int64, error) {
	if num > math.MaxInt64 {
		return 0, fmt.Errorf("uint64 %v can't tranfer to int64", num)
	}
	return int64(num), nil
}
