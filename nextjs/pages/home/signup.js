import React from 'react'
import { useForm } from 'react-hook-form'
import { signupFormOptions } from '../../app/config/formValidationSchema'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import { Button, Grid } from '@mui/material'
import { FormGroup, FormControlLabel, Checkbox, FormHelperText } from '@mui/material'
import theme from '../../src/theme'

function Signup() {

    const { register, handleSubmit, formState } = useForm(signupFormOptions)
    const { errors } = formState

    function onSubmit(data) {
        console.log(data)
        // do api stuff
    }


    let firstNameProps = {}
    if (errors.firstName) {
        firstNameProps.error = true
        firstNameProps.helperText = errors.firstName?.message
    }

    let lastNameProps = {}
    if (errors.lastName) {
        lastNameProps.error = true
        lastNameProps.helperText = errors.lastName?.message
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

    let confirmPasswordProps = {}
    if (errors.confirmPassword) {
        confirmPasswordProps.error = true
        confirmPasswordProps.helperText = errors.confirmPassword?.message
    }



    return (
        <Box
            sx={{
                my: 4,
                backgroundColor: 'secondary.light',

            }}
        >
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
                        <h1>Sign Up Form</h1>
                    </Grid>
                    <Grid item xs={6}>
                        <TextField
                            {...register("firstName")}
                            label="First Name"
                            variant="standard"
                            helperText="Please enter your first name"
                            {...firstNameProps}
                        />
                    </Grid>
                    <Grid item xs={6}>
                        <TextField
                            {...register("lastName")}
                            label="Last Name"
                            variant="standard"
                            helperText="Please enter your last name"
                            {...lastNameProps}
                        />
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
                    <Grid item xs={12}>
                        <TextField
                            {...register("confirmPassword")}
                            label="Confirm Password"
                            variant="standard"
                            helperText="Please enter your password again"
                            {...confirmPasswordProps}
                            fullWidth
                        />
                    </Grid>
                    <Grid item xs={6}>
                        <FormControlLabel
                            {...register("acceptTerms")}
                            control={<Checkbox value={true} />}
                            label="Accept Terms and Conditions"
                        />
                    </Grid>
                    <Grid item xs={6}>
                        {
                            errors.acceptTerms && <FormHelperText sx={{ color: theme.palette.error.main }}>{errors.acceptTerms?.message}</FormHelperText>
                        }
                    </Grid>

                    <Grid item xs={12} >
                        <Button type="submit" variant="contained">Submit</Button>
                    </Grid>
                </Grid>










            </form>
        </Box>
    )
}

// Signup.layout = "Fullpage"
export default Signup