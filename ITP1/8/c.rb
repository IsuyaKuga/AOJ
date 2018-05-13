h=Hash.new(0)
$<.read.downcase.chars{|c|h[c]+=1}
('a'..'z').each{|c|puts c+' : '+h[c].to_s}
