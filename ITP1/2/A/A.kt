fun main(args: Array<String>) {
  val line = readLine()!!.split(' ').map{x->x.toInt()}
  val a    = line[0]
  val b    = line[1]
  val result = if (a < b) {
    "a < b"
  } else if (a == b) {
    "a == b"
  } else {
    "a > b"
  }

  println("${result}")
}
