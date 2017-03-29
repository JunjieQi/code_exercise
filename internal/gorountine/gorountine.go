package gorountine

const max = 10

var Sem = make(chan bool, max)
