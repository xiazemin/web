package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"encoding/hex"
)
// 1 -- > 1
// 10-- > a
// 61-- > Z
const charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 将十进制转换为62进制  0-9a-zA-Z 六十二进制

//十进制转换成62进制
func base10ToBase62(n int64) string {
	var str string
	for n != 0 {
		str += string(charset[n % 62])
		n /= 62
	}

	for len(str) != 5 {
		str += "0"
	}
	return str
}


// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "00000000"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	if len(s)<8{
		b:=make([]byte,8-len(s))
		for i:=0;i<len(b);i++{
			b[i]='0'
		}
		s=string(b)+s
	}
	return s

}


//返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func transTo62(url string,key string)[]string {
	data := []byte(key+url)
	digest:=fmt.Sprintf("%x",md5.Sum(data))  //以十六进制数形式输出整数
	fmt.Println("digest length",len(digest))
	//128 bit,
	fmt.Println(len(md5.Sum(data)),md5.Sum(data),len(fmt.Sprintf("%b",md5.Sum(data))))
	//md5.Sum(data) 用10进制表示,只有16 位 16byte
	//将每一个unicode 转化成16 进制就变成了32位

	//16进制4 bit可以表示, unicode 8 bit 表示的

	var digestBin string

	for _,by:=range digest{ //16个word,32个16进制数,长度是32
		digestBin+=convertToBin(int(by))
	}

	//将长网址 md5 生成 32 位签名串,分为 4 段, 每段 8 个字节
	segments:=make([]string,4)
	fmt.Println(len(digest),digest,"\n",len(digestBin),digestBin)
	for i:=0;i<4;i++{
		segments[i]=string(digest[i*8:(i+1)*8])
		//hexData, _ := hex.DecodeString(segments[i])

		/*
		参数1 数字的字符串形式
		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
		 */
		hexData,e:=strconv.ParseInt(segments[i],16,64)
		fmt.Println(hexData,e)
		//对这四段循环处理, 取 8 个字节, 将他看成 16 进制串与 0x3fffffff(30位1) 与操作, 即超过 30 位的忽略处理
		hexData=hexData&0x3fffffff
		fmt.Println(hexData)

		//这 30 位分成 6 段, 每 5 位的数字作为字母表的索引取得特定字符, 依次进行获得 6 位字符串
		segments[i]=""
		for j:=0;j<6;j++{
			// 2^6=64 -2 =62 ,所以不是最末5位,是最末6位,弃掉第二位。不是0x0000001f,而是0x0000003d
			seg:=hexData&0x0000003d
			fmt.Println(seg)
			segments[i]+=string([]byte{charset[seg]})
			hexData=hexData>>5
		}
	}
	fmt.Println(segments,len(convertToBin(0x3fffffff)),convertToBin(0x0000003d))


	//总的 md5 串可以获得 4 个 6 位串,取里面的任意一个就可作为这个长 url 的短 url 地址

return segments

}

func main()  {
	s := fmt.Sprintf("%x",transTo62("http://www.me3.cn",""))
	fmt.Println(s)
	key:= "Leejor";
	fmt.Println(transTo62("http://www.me3.cn",key))
	//[fAVfui 3ayQry UZzyUr 36rQZn]
	fmt.Println(transTo62("http://www.me3.cn",""))
}

/*
ShortUrl(http://www.me3.cn")[0];  //得到值fAVfui

ShortUrl("http://www.me3.cn")[1];  //得到值3ayQry

ShortUrl("http://www.me3.cn")[2];  //得到值UZzyUr

ShortUrl("http://www.me3.cn")[3];  //得到值36rQZn
 */


/*
将长网址md5生成32位签名串，分为4段，每段8个字节；
对这四段循环处理，取8个字节，将他看成16进制串与0x3fffffff(30位1)与操作，即超过30位的忽略处理；
这30位分成6段，每5位的数字作为字母表的索引取得特定字符，依次进行获得6位字符串；
总的md5串可以获得4个6位串；取里面的任意一个就可作为这个长url的短url地址；
很简单的理论，我们并不一定说得到的URL是唯一的，但是我们能够取出4组URL，这样几乎不会出现太大的重复。
*/

/**
2进制 128 位  bit
16进制 32 位    每位4bit
4个部分,每个部分8个16进制数,占32bit
取 低30 bit
分6 部分,每部分 5bit 由于pow(2,6)-1 =64 ,目标是 62进制,所以,去掉倒数第二位
得到6个字母,  所以用到的只有md5 的1/4信息,可以有四个,如果冲突,可以继续使用,4个都冲突的概率很小了
 */