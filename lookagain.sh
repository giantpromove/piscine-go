find -depth -name "*.sh" | cut -d '.' -f2 | rev | cut -d "/" -f1| rev
