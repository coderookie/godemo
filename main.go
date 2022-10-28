package main

import (
	"godemo/models/protobuf"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
	_ "net/http/pprof"
	"flag"
)

func main() {
	// pprof性能监控, go run main.go --cpu_pprof_file=cpu.pprof, go tool pprof cpu.pprof
	var cpuPprofFile = flag.String("cpu_pprof_file", "", "write cpu info to file")
	var memPprofFile = flag.String("mem_pprof_file", "", "write mem info to file")
	flag.Parse()
	if *cpuPprofFile != ""{
		fp, err := os.Create(*cpuPprofFile)
		if err != nil {
			fmt.Printf("section[writeCpuToFile] func[os.Create] err[%s]\n", err.Error())
			os.Exit(1)
		}
		pprof.StartCPUProfile(fp)
		defer pprof.StopCPUProfile()
	}
	if *memPprofFile != "" {
		fp, err := os.Create(*memPprofFile)
		if err != nil {
			fmt.Printf("section[writeMemToFile] func[os.Create] err[%s]\n", err.Error())
			os.Exit(1)
		}
		pprof.WriteHeapProfile(fp)
		defer fp.Close()
	}

	// protobuf test start
	now1 := time.Now()
	data, err := user.GenerateUser2Protobuf()
	if err != nil {
		fmt.Printf("section[protobuf test generate], func[user.GenerateUser2Protobuf] err[%s]", err.Error())
		os.Exit(1)
	}
	sub1 := time.Now().Sub(now1).Microseconds()

	now2 := time.Now()
	zhanglei, err := user.ParseProtobuf2User(data)
	if err != nil {
		fmt.Println("section[protobuf test parse], func[user.ParseProtobuf2User] err[%s]", err.Error())
		os.Exit(2)
	}
	sub2 := time.Now().Sub(now2).Microseconds()

	now3 := time.Now()
	userJson, err := user.ParseUser2Json(zhanglei)
	if err != nil {
		fmt.Println("section[protobuf test parse2json] func[user.ParseUser2Json] err[%s]", err.Error())
		os.Exit(2)
	}
	sub3 := time.Now().Sub(now3).Microseconds()

	total := time.Now().Sub(now1).Microseconds()

	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("date[%s] section[protobuf test] unit[microsecond] total[%d] generateUser2ProtoCost[%d] parseProto2UserCost[%d] parseUser2Json[%d] output[%s]\n",
		date, total, sub1, sub2, sub3, userJson)
	fmt.Println()
	// protobuf test end

	fmt.Println("main func Done")
}
