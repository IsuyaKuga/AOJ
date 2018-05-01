gets
puts (m=->x{x.map{|l|l<<l.reduce(:+)}.transpose})[m[$<.map{|l|l.split.map &:to_i}]].map{|l|l*' '}
