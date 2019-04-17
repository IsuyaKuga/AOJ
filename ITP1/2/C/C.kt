fun main(argv: Array<String>) {
  val line = readLine()!!.split(' ').map{x->x.toInt()}
  var a    = line[0]
  var b    = line[1]
  var c    = line[2]

  if (c < b) {
    b = c.also{c = b}
  }
  if (b < a) {
    b = a.also{a = b}
  }
  if (c < b) {
    b = c.also{c = b}
  }

  println("${a} ${b} ${c}")
}
