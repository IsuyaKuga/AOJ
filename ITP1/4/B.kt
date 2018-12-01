fun main(argv: Array<String>) {
  val r = readLine()!!.toDouble()
  println("%.6f %.6f".format(Math.PI*r*r, 2*Math.PI*r))
}
