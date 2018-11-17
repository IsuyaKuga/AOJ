fun main(args: Array<String>) {
  val line = readLine()!!.split(' ').map{x->x.toInt()}
  val a    = line[0]
  val b    = line[1]
  val c    = line[2]

  val result = if (a < b && b < c) {
    "Yes"
  } else {
    "No"
  }
  println("${result}")
}
