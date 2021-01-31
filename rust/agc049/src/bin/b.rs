#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use std::usize::MAX;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
        s:Chars,
        t:Chars,
    }

    let mut ans = 0;
    let mut q = VecDeque::new();
    let mut p = MAX;
    for i in 0..n {
        if t[i] == '1' {
            q.push_back(i);
        }
        if s[i] == '1' {
            if p != MAX {
                ans += i - p;
                p = MAX;
            } else if let Some(v) = q.pop_front() {
                ans += i - v;
            } else {
                p = i;
            }
        }
    }

    if p == MAX && q.is_empty() {
        println!("{}", ans);
    } else {
        println!("{}", -1);
    }
}
