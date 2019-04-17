fun main(argv: Array<String>) {
  var count = 0

  val inputs = readLine()!!.split(' ').map{x->x.toInt()}
  val a = inputs[0]
  val b = inputs[1]
  val c = inputs[2]

  for (i in a..b) {
    if (c % i == 0) {
      count++
    }
  }
  println("${count}")
}
