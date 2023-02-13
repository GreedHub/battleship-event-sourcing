import React, { ReactElement } from "react";
import "./styles.scss";

type BoardParams = {
  className?: string;
};

export default function Board(params: BoardParams): ReactElement {
  const { className } = params;
  return <div className={`board ${className || ""}`}></div>;
}
