import React, { useContext } from 'react'
import { useForm } from 'react-hook-form'
import { loginFormOptions } from '../../config/formValidationSchema'
import TextField from '@mui/material/TextField'
import { Button, Grid } from '@mui/material'
import AuthContext from '../../api/auth/authContext'

function LoginForm() {

    const { register, handleSubmit, formState } = useForm(loginFormOptions)
    const { errors } = formState

    const { login, error, user, isLoading } = useContext(AuthContext)

    async function onSubmit({ email, password }) {
        login({ email, password })
    }

    let emailProps = {}
    if (errors.email) {
        emailProps.error = true
        emailProps.helperText = errors.email?.message
    }

    let passwordProps = {}
    if (errors.password) {
        passwordProps.error = true
        passwordProps.helperText = errors.password?.message
    }

    return (

        <form onSubmit={handleSubmit(onSubmit)} >
            <Grid
                container
                direction="row"
                rowSpacing={2}
                columnSpacing={1}
                maxWidth='sm'
                margin='auto'
            >
                <Grid item xs={12}>
                    <h1>Login Form</h1>
                </Grid>

                <Grid item xs={12}>
                    <TextField
                        {...register("email")}
                        label="Email"
                        variant="standard"
                        helperText="Please enter your email"
                        {...emailProps}
                        fullWidth
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        {...register("password")}
                        label="Password"
                        variant="standard"
                        helperText="Please enter your password"
                        {...passwordProps}
                        fullWidth
                    />
                </Grid>
                <Grid item xs={12} >
                    <Button type="submit" variant="contained">Submit</Button>
                </Grid>
            </Grid>


        </form>

    )
}

export default LoginForm