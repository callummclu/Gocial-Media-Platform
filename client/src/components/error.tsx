import { Title, Text, Card, Button } from "@mantine/core"
import styled from "styled-components"
import ReactMarkdown from 'react-markdown'
import dayjs from "dayjs"
import relativeTime from "dayjs/plugin/relativeTime"
import {removePost} from '../helpers/postHelper'
import { BiError } from "react-icons/bi"


export const Error = (props:any) => {

    return (
        <>
            <BiError/>
            <Title>404</Title>
            <Text>Page doesn't exist</Text>
        </>
    )
}
