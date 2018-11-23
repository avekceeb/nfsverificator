program summa;

Uses sysutils;

var hh, mm, ss: longint;
var time : string;

begin
  assign(input, 'input.txt');
  reset(input);
  read(time);
  writeln('t=', time);
  hh := StrToInt(copy(time, 1,2));
  mm := StrToInt(copy(time, 4,2));
  ss := StrToInt(copy(time, 7,2));
  writeln('h=', hh);
  writeln('m=', mm);
  writeln('s=', ss);
end.
