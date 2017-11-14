N = gets.strip.to_i
A = gets.split.map(&:to_i)
res = Array.new
res << A.join(' ') + "\n"
for i in 1..(N-1)
  key=A[i]
  j=i-1
  while j >= 0 and A[j] > key
    A[j+1]=A[j]
    j=j-1
  end
  A[j+1] = key
  res << A.join(' ') + "\n"
end
print res.join
