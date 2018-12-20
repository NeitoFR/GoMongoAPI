var a = "azer";
try {
  JSON.parse(a);
} catch (error) {
  console.log(error);
}
console.log("continue code");
