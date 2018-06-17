puts $<.read.chop.chop.split(/([a-z]+[\n0-9]+)/).reject(&:empty?).map(&:split).map{|a|w=a.shift;a.shift;(w*2)[a.map(&:to_i).reduce(:+).to_i%w.length,w.length]}
