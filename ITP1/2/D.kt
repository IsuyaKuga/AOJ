fun main(argv: Array<String>) {
  val inputs = readLine()!!.split(' ').map{x->x.toInt()}
  val W = inputs[0]
  val H = inputs[1]
  val x = inputs[2]
  val y = inputs[3]
  val r = inputs[4]

  if ((0 <= x-r) && (x+r <= W) && (0 <= y-r) && (y+r <= H)) {
    println("Yes")
  } else {
    println("No")
  }
}
