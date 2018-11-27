fun main(argv: Array<String>) {
  val inputs = readLine()!!.split(' ').map{x->x.toInt()}
  val a = inputs[0]
  val b = inputs[1]

  println("${a/b} ${a%b} %.5f".format(a.toDouble()/b.toDouble()))
}
