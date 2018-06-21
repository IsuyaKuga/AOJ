s=gets.chop
gets.to_i.times do
c,a,b,p=gets.split
a=a.to_i
b=b.to_i
case c
when 'print'
puts s[a..b]
when 'reverse'
s[a..b]=s[a..b].reverse
when 'replace'
s[a..b]=p
end
end
