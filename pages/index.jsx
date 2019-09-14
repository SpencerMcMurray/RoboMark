import React from "react";
import TestButtons from "../components/TestButtons";
import Layout from "../components/Layout";
import { Animated } from "react-animated-css";

const Index = props => {
  return (
    <Layout>
      <div
        style={{
          width: "auto",
          height: "50vh",
          background:
            "linear-gradient(90deg, rgba(131,58,180,1) 0%, rgba(253,29,29,1) 50%, rgba(252,176,69,1) 100%)"
        }}
      >
          <div className="row h-100 justify-content-center align-items-center">
            <div>
            <Animated
              animationIn="jackInTheBox"
              isVisible={true}
              className="fast"
            >
              <img
                src="../static/robomarklogo1.png"
                className="img-fluid rounded mx-auto d-block"
                alt="Responsive image"
                height="200px"
                width="200px"
                class="hover-grow"
                
              />
            </Animated>
            <h1 class="title-font hover-grow" className="mx-auto d-block"><b>RoboMark</b></h1>
            </div>
          </div>
      </div>
    </Layout>
  );
};

export default Index;
