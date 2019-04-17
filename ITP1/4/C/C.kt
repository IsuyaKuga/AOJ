fun main(argv: Array<String>) {
  val list: MutableList<Int> = mutableListOf()

  loop@ while (true) {
    val inputs = readLine()!!.split(' ')
    val a  = inputs[0].toInt()
    val op = inputs[1]
    val b  = inputs[2].toInt()
    val res = when (op) {
      "+"  -> a+b
      "-"  -> a-b
      "*"  -> a*b
      "/"  -> a/b
      else -> break@loop
    }
    list.add(res)
  }

  for (i in list.indices) {
    println("${list[i]}")
  }
}
