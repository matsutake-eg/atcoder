#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::{fastout, input, marker::*};
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: don't print in 'match arm' because of fastout's bug
#[fastout]
fn main() {
    input! {
        n:usize,
        s:Chars,
    }

    let mut dp = vec![vec![mod_int(0); 8]; n + 1];
    dp[0][0] = mod_int(1);
    for i in 0..n {
        for j in 0..8 {
            dp[i + 1][j] = dp[i + 1][j] + dp[i][j];
            if (s[i] == 'a' && j == 0)
                || (s[i] == 't' && j == 1)
                || (s[i] == 'c' && j == 2)
                || (s[i] == 'o' && j == 3)
                || (s[i] == 'd' && j == 4)
                || (s[i] == 'e' && j == 5)
                || (s[i] == 'r' && j == 6)
            {
                dp[i + 1][j + 1] = dp[i + 1][j + 1] + dp[i][j];
            }
        }
    }

    println!("{}", dp[n][7].value());
}

use std::ops::{
    Add, AddAssign, BitAnd, Div, DivAssign, Mul, MulAssign, RemAssign, ShrAssign, Sub, SubAssign,
};

fn mod_int(v: usize) -> ModInt<usize> {
    ModInt::new(v, 1_000_000_007)
}
pub struct ModInt<T> {
    v: T,
    m: T,
}

impl<T> ModInt<T>
where
    T: Copy,
{
    pub fn value(&self) -> T {
        self.v
    }
    pub fn modulo(&self) -> T {
        self.m
    }
}

impl<T> ModInt<T> {
    fn new_unchecked(v: T, modulo: T) -> Self {
        Self { v, m: modulo }
    }
}

impl<T> ModInt<T>
where
    T: Copy + RemAssign + PartialOrd,
{
    pub fn new(mut v: T, modulo: T) -> Self {
        if v >= modulo {
            v %= modulo;
        }
        Self::new_unchecked(v, modulo)
    }
}

impl<T> ModInt<T>
where
    T: Copy
        + Sub<Output = T>
        + ShrAssign
        + BitAnd<Output = T>
        + PartialEq
        + PartialOrd
        + Div<Output = T>
        + RemAssign,
    ModInt<T>: MulAssign,
{
    pub fn pow(self, e: T) -> Self {
        let zero = self.modulo() - self.modulo();
        let one = self.modulo() / self.modulo();
        let mut e = e;
        let mut result = Self::new_unchecked(one, self.modulo());
        let mut cur = self;
        while e > zero {
            if e & one == one {
                result *= cur;
            }
            e >>= one;
            cur *= cur;
        }
        result
    }
}

impl<T> Copy for ModInt<T> where T: Copy {}
impl<T> Clone for ModInt<T>
where
    T: Copy,
{
    fn clone(&self) -> Self {
        Self::new_unchecked(self.value(), self.modulo())
    }
}

impl<T> Add<T> for ModInt<T>
where
    T: AddAssign + SubAssign + RemAssign + Copy + PartialOrd,
{
    type Output = Self;
    fn add(self, mut rhs: T) -> Self::Output {
        if rhs >= self.modulo() {
            rhs %= self.modulo();
        }
        rhs += self.value();
        if rhs >= self.modulo() {
            rhs -= self.modulo();
        }
        Self::new_unchecked(rhs, self.modulo())
    }
}

impl<T> Sub<T> for ModInt<T>
where
    T: AddAssign + SubAssign + RemAssign + Copy + PartialOrd,
{
    type Output = Self;
    fn sub(self, mut rhs: T) -> Self::Output {
        if rhs >= self.modulo() {
            rhs %= self.modulo();
        }

        let mut result = self.value();
        result += self.modulo();
        result -= rhs;

        if result >= self.modulo() {
            result -= self.modulo();
        }
        Self::new_unchecked(result, self.modulo())
    }
}

impl<T> Mul<T> for ModInt<T>
where
    T: MulAssign + RemAssign + Copy + PartialOrd,
{
    type Output = Self;
    fn mul(self, mut rhs: T) -> Self::Output {
        if rhs >= self.modulo() {
            rhs %= self.modulo();
        }
        rhs *= self.value();
        rhs %= self.modulo();
        Self::new_unchecked(rhs, self.modulo())
    }
}

impl<T> Add<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Add<T, Output = ModInt<T>>,
{
    type Output = Self;
    fn add(self, rhs: ModInt<T>) -> Self::Output {
        self + rhs.value()
    }
}
impl<T> Sub<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Sub<T, Output = ModInt<T>>,
{
    type Output = Self;
    fn sub(self, rhs: ModInt<T>) -> Self::Output {
        self - rhs.value()
    }
}
impl<T> Mul<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Mul<T, Output = ModInt<T>>,
{
    type Output = Self;
    fn mul(self, rhs: ModInt<T>) -> Self::Output {
        self * rhs.value()
    }
}
impl<T> Div<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Div<T, Output = ModInt<T>>,
{
    type Output = Self;
    fn div(self, rhs: ModInt<T>) -> Self::Output {
        self / rhs.value()
    }
}

impl<T> AddAssign<T> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Add<T, Output = ModInt<T>>,
{
    fn add_assign(&mut self, other: T) {
        *self = *self + other;
    }
}
impl<T> AddAssign<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Add<ModInt<T>, Output = ModInt<T>>,
{
    fn add_assign(&mut self, other: ModInt<T>) {
        *self = *self + other;
    }
}

impl<T> SubAssign<T> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Sub<T, Output = ModInt<T>>,
{
    fn sub_assign(&mut self, other: T) {
        *self = *self - other;
    }
}

impl<T> SubAssign<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Sub<ModInt<T>, Output = ModInt<T>>,
{
    fn sub_assign(&mut self, other: ModInt<T>) {
        *self = *self - other;
    }
}

impl<T> DivAssign<T> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Div<T, Output = ModInt<T>>,
{
    fn div_assign(&mut self, rhs: T) {
        *self = *self / rhs
    }
}
impl<T> DivAssign<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Div<ModInt<T>, Output = ModInt<T>>,
{
    fn div_assign(&mut self, rhs: ModInt<T>) {
        *self = *self / rhs
    }
}

impl<T> MulAssign<T> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Mul<T, Output = ModInt<T>>,
{
    fn mul_assign(&mut self, rhs: T) {
        *self = *self * rhs;
    }
}

impl<T> MulAssign<ModInt<T>> for ModInt<T>
where
    T: Copy,
    ModInt<T>: Mul<ModInt<T>, Output = ModInt<T>>,
{
    fn mul_assign(&mut self, rhs: ModInt<T>) {
        *self = *self * rhs;
    }
}

impl<T> Div<T> for ModInt<T>
where
    T: Copy
        + Add<Output = T>
        + Sub<Output = T>
        + Div<Output = T>
        + BitAnd<Output = T>
        + PartialEq
        + PartialOrd
        + ShrAssign
        + RemAssign
        + MulAssign,
{
    type Output = Self;
    fn div(self, mut rhs: T) -> Self::Output {
        if rhs >= self.modulo() {
            rhs %= self.modulo();
        }
        let one = self.modulo() / self.modulo();
        let two = one + one;
        self * Self::new_unchecked(rhs, self.modulo()).pow(self.modulo() - two)
    }
}
