gets;puts$<.read.split(/\n/).map(&:split).map{|t,h|t==h ?[1,1]:t<h ?[0,3]:[3,0]}.reduce([0,0]){|s,a|[s[0]+a[0],s[1]+a[1]]}.join(' ')
