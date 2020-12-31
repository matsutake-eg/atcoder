use proconio;

#[allow(unused_imports)]
use proconio::marker::*;
#[allow(unused_imports)]
use std::cmp::*;
#[allow(unused_imports)]
use std::collections::*;
#[allow(unused_imports)]
use std::f64::consts::*;

#[allow(unused)]
const INF: usize = std::usize::MAX / 4;
#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        s:Chars,
        t:Chars,
    }

    let mut ans = INF;
    for i in 0..s.len() - t.len() + 1 {
        let mut cnt = 0;
        for j in 0..t.len() {
            if s[i + j] != t[j] {
                cnt += 1;
            }
        }
        ans = min(ans, cnt);
    }
    println!("{}", ans);
}
