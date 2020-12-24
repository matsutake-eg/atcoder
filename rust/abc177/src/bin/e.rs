#![allow(unused_imports)]
use num_integer::gcd;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;

#[allow(unused)]
const INF: usize = std::usize::MAX / 4;
#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[fastout]
fn main() {
    input! {
        n:usize,
        a:[usize;n],
    }

    if a.iter().fold(a[0], |acc, &x| gcd(acc, x)) != 1 {
        println!("not coprime");
        return;
    }

    let mut hm = HashSet::new();
    for v in a {
        for pv in prime_fac(v) {
            if hm.contains(&pv) {
                println!("setwise coprime");
                return;
            }
            hm.insert(pv);
        }
    }
    println!("pairwise coprime");
}

fn prime_fac(mut x: usize) -> HashSet<usize> {
    let mut pfs = HashSet::new();
    for i in 2..x {
        if i * i > x {
            break;
        }
        while x % i == 0 {
            pfs.insert(i);
            x /= i;
        }
    }
    pfs.insert(x);
    pfs.remove(&1);
    pfs
}
