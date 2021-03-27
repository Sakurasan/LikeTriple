package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	mid  string
	bv   string
	_url = "http://api.bilibili.com/x/space/coin/video?vmid="

	hookableSignals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	}
	defaultHeartbeatTime = 30 * time.Second //1 * time.Minute
	wg                   sync.WaitGroup
)

type coinVideo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    []data `json:"data"`
}

type data struct {
	Aid       int    `json:"aid"`
	Videos    int    `json:"videos"`
	Tid       int    `json:"tid"`
	Tname     string `json:"tname"`
	Copyright int    `json:"copyright"`
	Pic       string `json:"pic"`
	Title     string `json:"title"`
	Pubdate   int    `json:"pubdate"`
	Ctime     int    `json:"ctime"`
	Desc      string `json:"desc"`
	State     int    `json:"state"`
	Duration  int    `json:"duration"`
	MissionID int    `json:"mission_id,omitempty"`
	Rights    struct {
		Bp            int `json:"bp"`
		Elec          int `json:"elec"`
		Download      int `json:"download"`
		Movie         int `json:"movie"`
		Pay           int `json:"pay"`
		Hd5           int `json:"hd5"`
		NoReprint     int `json:"no_reprint"`
		Autoplay      int `json:"autoplay"`
		UgcPay        int `json:"ugc_pay"`
		IsCooperation int `json:"is_cooperation"`
		UgcPayPreview int `json:"ugc_pay_preview"`
		NoBackground  int `json:"no_background"`
	} `json:"rights"`
	Owner struct {
		Mid  int    `json:"mid"`
		Name string `json:"name"`
		Face string `json:"face"`
	} `json:"owner"`
	Stat struct {
		Aid      int `json:"aid"`
		View     int `json:"view"`
		Danmaku  int `json:"danmaku"`
		Reply    int `json:"reply"`
		Favorite int `json:"favorite"`
		Coin     int `json:"coin"`
		Share    int `json:"share"`
		NowRank  int `json:"now_rank"`
		HisRank  int `json:"his_rank"`
		Like     int `json:"like"`
		Dislike  int `json:"dislike"`
	} `json:"stat"`
	Dynamic   string `json:"dynamic"`
	Cid       int    `json:"cid"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
	ShortLink  string `json:"short_link"`
	Bvid       string `json:"bvid"`
	Coins      int    `json:"coins"`
	Time       int    `json:"time"`
	IP         string `json:"ip"`
	InterVideo bool   `json:"inter_video"`
}

func handleSignal() {
	pid := syscall.Getpid()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, hookableSignals...)

	go func() {
		sig := <-sigs
		log.Printf("\npid[%d], signal: [%v]", pid, sig)
		done <- true
	}()

	for {
		select {
		case <-done:
			log.Println("have done")
			return
		default:
			<-time.After(defaultHeartbeatTime)
		}
	}

}

func getCoinVideo(vmid string) (cv *coinVideo, err error) {
	req, err := http.NewRequest("GET", _url+vmid, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(resp.Body).Decode(&cv); err != nil {
		return
	}
	return
}

func init() {
	if os.Getenv("mid") == "" {
		panic("mid not set!")
	}
	mid = os.Getenv("mid")
}

func main() {
	go func() {
		log.Println("LikeTriple(一键三连) 启动成功 \n\n哔哩哔哩 (゜-゜)つロ 干杯~-bilibili\n ")
		for range time.Tick(defaultHeartbeatTime) {
			video, err := getCoinVideo(mid) //19161224
			if err != nil {
				log.Println("getCoinVideo Err:", err.Error())
				continue
			}
			if bv == "" {
				if fileISexist(video.Data[0].Title + ".mp4") {
					log.Println("[文件已存在]", video.Data[0].Title+".mp4")
					bv = video.Data[0].Bvid
					continue
				}
				if fileISexist(video.Data[0].Title + ".flv") {
					log.Println("[文件已存在]", video.Data[0].Title+".flv")
					bv = video.Data[0].Bvid
					continue
				}
				bv = video.Data[0].Bvid
				log.Println(fmt.Sprintf("首次下载 av%d", video.Data[0].Aid))
			} else if bv == video.Data[0].Bvid {
				continue
			} else {
				log.Println(video.Data[0].Title)
				log.Println(video.Data[0].ShortLink)
				bv = video.Data[0].Bvid
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				now := time.Now()
				var cmd *exec.Cmd
				if os.Getenv("format") == "flv" {
					cmd = exec.Command("you-get", "-o", "/download", video.Data[0].ShortLink)
				} else {
					cmd = exec.Command("you-get", "-o", "/download", "--format=dash-flv", video.Data[0].ShortLink)
				}

				out, err := cmd.Output()
				if err != nil {
					log.Println("getCoinVideo Err:", err.Error())
					return
				}
				since := time.Since(now).Seconds()
				if since/float64(60) > 1 {
					log.Println("\n"+string(out), fmt.Sprintf("耗时%.2f分钟\n--------------", since/float64(60)))
				} else {
					log.Println("\n"+string(out), fmt.Sprintf("耗时%.f秒\n--------------", since))
				}
			}()
			wg.Wait()
			continue
		}
	}()
	handleSignal()
}

func fileISexist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
