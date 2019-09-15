import React from "react";
import Layout from "../components/Layout";
import { Animated } from "react-animated-css";
import { Button, Card } from "react-bootstrap";
import Link from "next/link";

const userId = 1;

const Index = props => {
  return (
    <Layout>
      <div
        style={{
          width: "auto",
          height: "40vh",
          background:
            "linear-gradient(90deg, rgba(131,58,180,1) 0%, rgba(253,29,29,1) 50%, rgba(252,176,69,1) 100%)"
        }}
      >
        <div className="h-100 justify-content-center align-items-center pt-4">
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
            />
          </Animated>
          <h1 className="mx-auto d-block title-font hover-grow text-center">
            <b>RoboMark</b>
          </h1>
        </div>
      </div>
      <div
        style={{
          width: "auto",
          height: "60vh",
          background: "#f8f8f8"
        }}
      >
        <Card className="text-center">
          <Card.Header>Welcome!</Card.Header>
          <Card.Body className="p-5">
            <div>
              <Card.Title>What is RoboMark?</Card.Title>
              <Card.Text>
                RoboMark allows teachers to automate the process of
                marking,uploading and distributing tests.
                <br></br>RoboMark is built using React, Azure, Python and Go.
              </Card.Text>
              <Link href={`/tests?user=${userId}`}>
                <Button variant="primary">Tests</Button>
              </Link>
            </div>
          </Card.Body>
          <Card.Footer className="text-muted">
            <a
              target="_blank"
              href="https://github.com/SpencerMcMurray/RoboMark"
            >
              <i
                className="fab fa-github"
                style={{
                  fontSize: "40px"
                }}
              ></i>
            </a>
          </Card.Footer>
        </Card>
      </div>
    </Layout>
  );
};

export default Index;
