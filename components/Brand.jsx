import React from "react";
import { Animated } from "react-animated-css";
import Link from "next/link";

const Brand = () => {
  return (
    <Animated animationIn="rollIn" isVisible={true}>
      <div className="container">
      <div className="row align-items-center">
      <img
        src="../static/robomarklogo1.png"
        className="img-fluid rounded m-1 float-left"
        alt="Responsive image"
        height="80px"
        width="80px"
      />
      <Link href="/">
        <h1 className="title-font text-center hover-grow float-right text-white"><b>RoboMark</b></h1>
      </Link>
      </div>
      </div>
    </Animated>
  );
};

export default Brand;
