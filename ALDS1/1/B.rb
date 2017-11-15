def gcd(x, y)
  if x < y
    x, y = y, x
  end
  return y if x % y == 0
  gcd y, x%y
end

x,y = gets.split.map(&:to_i)
puts gcd(x, y)
