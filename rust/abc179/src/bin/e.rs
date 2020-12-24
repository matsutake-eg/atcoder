#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        x:usize,
        m:usize,
    }

    let mut id = vec![std::usize::MAX; m];
    let mut x = x;
    let mut a = vec![];
    let mut len = 0;
    let mut total = 0;
    while id[x] == std::usize::MAX {
        a.push(x);
        id[x] = len;
        len += 1;
        total += x;
        x = x * x % m;
    }

    if n <= len {
        println!("{}", a[0..n].iter().fold(0, |acc, x| acc + x));
        return;
    }

    let mut ans = total;
    let id_last = id[x];
    let total_loop = a[id_last..len].iter().fold(0, |acc, x| acc + x);
    let mut n = n;
    n -= len;
    let c = len - id_last;
    ans += total_loop * (n / c);
    n %= c;
    ans += a[id_last..(id_last + n)].iter().fold(0, |acc, x| acc + x);
    println!("{}", ans);
}
