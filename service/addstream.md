12、/index/api/addStreamProxy
功能：动态添加 rtsp/rtmp/hls/http-ts/http-flv 拉流代理(只支持 H264/H265/aac/G711/opus 负载)

范例：http://127.0.0.1/index/api/addStreamProxy?vhost=__defaultVhost__&app=proxy&stream=0&url=rtmp://live.hkstv.hk.lxdns.com/live/hks2

参数：

参数	参数类型	释意	是否必选
secret	string	api 操作密钥(配置文件配置)	Y
vhost	string	添加的流的虚拟主机，例如__defaultVhost__	Y
app	string	添加的流的应用名，例如 live	Y
stream	string	添加的流的 id 名，例如 test	Y
url	string	拉流地址，例如 rtmp://live.hkstv.hk.lxdns.com/live/hks2	Y
retry_count	int	拉流重试次数，默认为-1 无限重试	N
rtp_type	int	rtsp 拉流时，拉流方式，0：tcp，1：udp，2：组播	N
timeout_sec	int	拉流超时时间，单位秒，float 类型	N
enable_hls	bool	是否转换成 hls-mpegts 协议	N
enable_hls_fmp4	bool	是否转换成 hls-fmp4 协议	N
enable_mp4	bool	是否允许 mp4 录制	N
enable_rtsp	bool	是否转 rtsp 协议	N
enable_rtmp	bool	是否转 rtmp/flv 协议	N
enable_ts	bool	是否转 http-ts/ws-ts 协议	N
enable_fmp4	bool	是否转 http-fmp4/ws-fmp4 协议	N
hls_demand	bool	该协议是否有人观看才生成	N
rtsp_demand	bool	该协议是否有人观看才生成	N
rtmp_demand	bool	该协议是否有人观看才生成	N
ts_demand	bool	该协议是否有人观看才生成	N
fmp4_demand	bool	该协议是否有人观看才生成	N
enable_audio	bool	转协议时是否开启音频	N
add_mute_audio	bool	转协议时，无音频是否添加静音 aac 音频	N
mp4_save_path	string	mp4 录制文件保存根目录，置空使用默认	N
mp4_max_second	int	mp4 录制切片大小，单位秒	N
mp4_as_player	bool	MP4 录制是否当作观看者参与播放人数计数	N
hls_save_path	string	hls 文件保存保存根目录，置空使用默认	N
modify_stamp	int	该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)	N
auto_close	bool	无人观看是否自动关闭流(不触发无人观看 hook)	N
响应：


{
   "code" : 0,
   "data" : {
      "key" : "__defaultVhost__/proxy/0"  # 流的唯一标识
   }
}
#