r,c = gets.split.map!(&:to_i)

res = String.new
raw_sum = Array.new(c+1,0)

r.times do
  line = gets.split.map!(&:to_i)
  i = 0
  line << line.inject(0) {|col_sum, element|
    raw_sum[i]+=element
    i += 1
    col_sum + element}
  raw_sum[c] += line[c]
  res += line.map(&:to_s).join(' ')+"\n"
end
print res
puts raw_sum.map(&:to_s).join(' ')
