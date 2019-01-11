package main

//func swap(a,b *int){
//	*b,*a=*a,*b
//}
//
//func lengthNonrepeatingSubStr(s string)int{
//	lastOccured :=make(map[byte]int)
//	start:=0
//	maxlength:=0
//
//	for i,ch:=range []byte(s){
//		lastI,ok:=lastOccured[ch]//lastI是遍历的字符的上一次出现过的下标，默认为0因为map的V默认没有内容，Ok代表V是否有内容，也就是是否出现过
//
//		if ok&&lastI>=start {
//			start=lastI+1//关键要理解这里的start即最长不重复字串的开头必须要在上一个出现了的字母的后面
//		}
//		if i-start+1>maxlength{
//			maxlength=i-start+1
//		}
//		lastOccured[ch]=i
//		fmt.Println(ch,lastI,start,ok,maxlength,lastOccured[ch])
//	}
//	return maxlength
//}
//func createTreeNode() *TreeNode{
//	return &TreeNode{Value:value}
//}
//type Retriever struct {
//	UserAgent string
//	TimeOut  time.Duration
//}
//func (r Retriever)Get(url string) string{
//	resp,err:=http.Get(url)
//	if err!=nil{panic(err)}
//	result,err:=httputil.DumpResponse(resp,true)
//	if err
//}
//func main() {
	//fmt.Println("Hello World!")
	//if content,err:=ioutil.ReadFile("xk.txt"); err==nil{
	//
	//	fmt.Println(string(content))
	//}else {
	//	fmt.Println(err)
	//}
	//a,b:=3,4
	//swap(&a,&b)
	//fmt.Println(a,b)
	//arr:=[...]int{0,1,2,3,4,5,6,7}
	//s1:=arr[2:6]
	//s2:=s1[3:5]
	//fmt.Println(s2)
	//fmt.Println(lengthNonrepeatingSubStr("baca"))
	//root.Left.Right=createTreeNode(2)
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)
	//bufio.NewReader(r).Peek(1024)
//}