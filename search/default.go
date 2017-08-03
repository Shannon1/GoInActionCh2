package search

// 可以将defaultMatcher理解为Matcher（接口）类型的一个实现
type defaultMatcher struct {}


// 程序里所有的 init 方法都会在 main 函数启动前被调用
// 导入包时的下划线告诉编译器找到并执行这个包里的init函数
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 如果声明函数的时候带有接收者，则意味着声明了一个方法。这个方法会和指定的接收者的类型(defaultMatcher)绑在一起。
// (可以理解为面向对象中一个类的成员函数)
// defaultMatcher实现了Matcher接口声明的Search方法，所以defaultMatcher类型就是一个Matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
