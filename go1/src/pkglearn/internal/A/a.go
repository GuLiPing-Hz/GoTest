/**内部包，只允许在同一个层级的允许访问，外部无法访问
Go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处 理。这种包叫internal包

一个internal包只能被和internal目录有同一个父目录的包所导入
*/
package A

const (
	A = 1024
	a = 2048
)
