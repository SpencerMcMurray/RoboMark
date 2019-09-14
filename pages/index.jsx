import React from "react";
import TestButtons from "../components/TestButtons";
import Layout from "../components/Layout";
import { Animated } from "react-animated-css";

const Index = props => {
  return (
    <Layout>
      <div class="jumbotron jumbotron-fluid ">
        <div class="container ">
        </div>
      </div>

      <Animated animationIn="jackInTheBox" isVisible={true} className="fast">
        <img
          src="../static/robomarklogo.png"
          class="img-fluid rounded mx-auto d-block"
          alt="Responsive image"
        />
      </Animated>
    </Layout>
  );
};

export default Index;
