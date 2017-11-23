def prime?(input)
  square_root = input.**(0.5).floor.to_i
  2.upto(square_root) do |divisor|
    return false if input % divisor == 0
  end
  return true
end

num = gets.chomp.to_i
prime_num = 0

num.times do
  input = gets.chomp.to_i
  prime_num += 1 if prime?(input)
end

puts prime_num
