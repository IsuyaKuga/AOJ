fun main(args: Array<String>) {
  val x = readLine()!!.split(' ').map{x->x.toInt()}
  println("${x[0]*x[1]} ${2*(x[0]+x[1])}")
}
