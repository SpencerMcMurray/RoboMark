import { Spinner } from "react-bootstrap";

const LoadingScreen = props => {
  return (
    <div
      style={{ height: "100vh" }}
      className="w-100 d-flex justify-content-center align-items-center"
    >
      <Spinner
        animation="border"
        variant={props.color ? props.color : "primary"}
      />
    </div>
  );
};

export default LoadingScreen;
